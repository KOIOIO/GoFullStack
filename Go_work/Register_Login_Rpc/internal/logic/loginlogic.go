package logic

import (
	"context"
	"errors"

	"Register_Login_Rpc/User"
	"Register_Login_Rpc/internal/svc"
	"Register_Login_Rpc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *User.LoginRequest) (*User.LoginResponse, error) {
	// todo: add your logic here and delete this line
	if in.Username == "" || in.Password == "" {
		return &User.LoginResponse{
			Code:    400,
			Message: "Username or password cannot be empty",
		}, errors.New("Username or password cannot be empty")
	}

	//从MongoDB获取用户信息
	user, err := l.GetFromMongo(in.Username)
	if err != nil {
		return &User.LoginResponse{
			Code:    400,
			Message: "Username is incorrect or not exit",
		}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		return &User.LoginResponse{
			Code:    400,
			Message: "Password is incorrect",
		}, nil
	}

	return &User.LoginResponse{
		Code:    200,
		Message: "Login success",
	}, nil
}

func (l *LoginLogic) GetFromMongo(Username string) (*types.UserInfo, error) {
	// 连接到MongoDB并获取用户信息的逻辑
	// 假设已经有一个MongoDB客户端实例 svcCtx.MongoDB
	collection := l.svcCtx.MongoDB.Database(l.svcCtx.Config.Mongo.Database).Collection("users")

	var user types.UserInfo
	err := collection.FindOne(l.ctx, bson.M{"username": Username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
