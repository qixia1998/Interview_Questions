package logic

import (
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/rpc/model"
	"context"
	"google.golang.org/grpc/status"

	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/rpc/courseware"
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *courseware.AddRequest) (*courseware.AddResponse, error) {
	_, err := l.svcCtx.CoursewareModel.FindOneByCode(l.ctx, in.Code)
	if err == nil {
		return nil, status.Error(5000, "课件已存在")
	}

	if err == model.ErrNotFound {
		newCourseWare := model.Courseware{
			Code: in.Code,
			Name: in.Name,
			Type: in.Type,
		}

		_, err = l.svcCtx.CoursewareModel.Insert(l.ctx, &newCourseWare)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &courseware.AddResponse{}, nil
	}

	return nil, status.Error(500, err.Error())
}
