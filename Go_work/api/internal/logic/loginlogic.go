package logic

import (
	"Register_Login_Rpc/User"
	"context"
	"strings"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginRequest) (resp *types.UserLoginResponse, err error) {
	// todo: add your logic here and delete this line
	username := strings.TrimSpace(req.Username)
	passwprd := strings.TrimSpace(req.Password)
	auth := l.svcCtx.Config.Auth

	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &User.LoginRequest{
		Username: username,
		Password: passwprd,
	})

	if loginResp.Code != 200 {
		return &types.UserLoginResponse{
			Code:    404,
			Message: "login failed",
			Token:   "",
		}, nil
	}

	token, err := GenToken(JwtPayLoad{
		Username: req.Username,
		Code:     200,
	}, auth.AccessSecret, auth.AccessExpire)

	if err != nil {
		return &types.UserLoginResponse{
			Code:    500,
			Message: "generate token failed",
			Token:   "",
		}, nil
	}

	return &types.UserLoginResponse{
		Code:    200,
		Message: "login success",
		Token:   token,
	}, nil
}

// GenToken generate token
