package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mongo struct {
		Hosts      []string
		Database   string
		User       string
		Password   string
		AuthSource string
	}
}
