package svc

import (
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/api/internal/config"
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/rpc/coursewareclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CourseWareRpc coursewareclient.Courseware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CourseWareRpc: coursewareclient.NewCourseware(zrpc.MustNewClient(c.CourseWareRpc)),
	}
}
