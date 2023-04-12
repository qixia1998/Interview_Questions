package logic

import (
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/rpc/coursewareclient"
	"context"

	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/api/internal/svc"
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoursewareUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoursewareUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoursewareUpdateLogic {
	return &CoursewareUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoursewareUpdateLogic) CoursewareUpdate(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	_, err = l.svcCtx.CourseWareRpc.Update(l.ctx, &coursewareclient.UpdateRequest{
		Id:   req.Id,
		Code: req.Code,
		Name: req.Name,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{}, nil
}
