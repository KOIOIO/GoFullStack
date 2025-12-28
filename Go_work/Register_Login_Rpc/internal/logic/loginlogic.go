package logic

import (
	"context"

	"Register_Login_Rpc/User"
	"Register_Login_Rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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

	return &User.LoginResponse{}, nil
}
