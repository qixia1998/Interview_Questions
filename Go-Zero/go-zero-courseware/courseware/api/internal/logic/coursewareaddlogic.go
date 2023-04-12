package logic

import (
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/rpc/coursewareclient"
	"context"

	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/api/internal/svc"
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoursewareAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoursewareAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoursewareAddLogic {
	return &CoursewareAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoursewareAddLogic) CoursewareAdd(req *types.AddRequest) (resp *types.AddResponse, err error) {
	_, err = l.svcCtx.CourseWareRpc.Add(l.ctx, &coursewareclient.AddRequest{
		Code: req.Code,
		Name: req.Name,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddResponse{}, nil
}
