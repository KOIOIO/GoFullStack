package logic

import (
	"Register_Login_Rpc/User"
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// todo: add your logic here and delete this line
	username := req.UserName
	password := req.Password
	confirmPassword := req.ConfirmPassword
	email := req.Email

	RegisterResponse, err := l.svcCtx.UserRpc.Register(l.ctx, &User.RegisterRequest{
		Username:        username,
		Password:        password,
		ConfirmPassword: confirmPassword,
		Email:           email,
	})

	if RegisterResponse.Code != 200 {
		return &types.UserRegisterResponse{
			Code:    1001,
			Message: "register failed",
			UserId:  "",
		}, nil
	}

	return &types.UserRegisterResponse{
		Code:    200,
		Message: "register success",
		UserId:  RegisterResponse.UserId,
	}, nil

}
