package svc

import (
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/rpc/internal/config"
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	CoursewareModel model.CoursewareModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		CoursewareModel: model.NewCoursewareModel(conn, c.CacheRedis),
	}
}
