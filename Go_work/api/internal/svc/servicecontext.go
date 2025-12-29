package svc

import (
	"api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userservice.UserServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
