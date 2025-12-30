package logic

import (
	"context"
	"fmt"
	"time"

	"Register_Login_Rpc/User"
	"Register_Login_Rpc/internal/svc"
	"Register_Login_Rpc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *User.RegisterRequest) (*User.RegisterResponse, error) {
	// Check if passwords match
	if in.Password != in.ConfirmPassword {
		return &User.RegisterResponse{
			Code:    400,
			Message: "Passwords do not match",
		}, nil
	}

	//todo 目前先使用逻辑查重，后面在学MongoDB是否可以建立唯一索引
	collection := l.svcCtx.MongoDB.Database(l.svcCtx.Config.Mongo.Database).Collection("users")
	count, err := collection.CountDocuments(l.ctx, bson.M{"username": in.Username})
	if err != nil {
		return &User.RegisterResponse{
			Code:    500,
			Message: "Error checking username duplication",
		}, err
	}

	if count > 0 {
		return &User.RegisterResponse{
			Code:    409,
			Message: "Username already exists",
		}, nil
	}

	EmailCount, err := l.svcCtx.MongoDB.Database(l.svcCtx.Config.Mongo.Database).Collection("users").CountDocuments(l.ctx, bson.M{"email": in.Email})
	if err != nil {
		return &User.RegisterResponse{
			Code:    500,
			Message: "Error checking email duplication",
		}, err
	}

	if EmailCount > 0 {
		return &User.RegisterResponse{
			Code:    409,
			Message: "Email already exists",
		}, nil
	}

	tempPassword, err := l.bcryptPassword(in.Password)
	if err != nil {
		return &User.RegisterResponse{
			Code:    500,
			Message: "Error encrypting password",
		}, err
	}
	// 承接用户
	newUser := types.UserInfo{
		Username:  in.Username,
		Password:  tempPassword, // 生产环境务必加密（如bcrypt）
		Email:     in.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert the new user
	result, err := collection.InsertOne(l.ctx, newUser)
	if err != nil {
		return &User.RegisterResponse{
			Code:    500,
			Message: "Failed to register user",
		}, err
	}

	if result.InsertedID == nil {
		return &User.RegisterResponse{
			Code:    500,
			Message: "Failed to register user",
		}, nil
	}
	//将InsertedID转为ObjectId字符串
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return &User.RegisterResponse{
			Code:    500,
			Message: "Failed to parse ObjectID",
		}, fmt.Errorf("invalid ObjectID type")
	}
	objectIDStr := insertedID.Hex()

	// 返回结果：新增_id字段
	return &User.RegisterResponse{
		Code:    200,
		Message: "User registered successfully",
		UserId:  objectIDStr,
	}, nil
}

func (l *RegisterLogic) bcryptPassword(password string) (string, error) {
	hashbytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashbytes), nil
}
