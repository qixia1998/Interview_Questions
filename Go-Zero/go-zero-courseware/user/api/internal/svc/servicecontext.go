package svc

import (
	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/config"
	"Interview_Questions/Go-Zero/go-zero-courseware/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc userclient.USer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUSer(zrpc.MustNewClient(c.UserRpc)),
	}
}
