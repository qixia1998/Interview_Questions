package logic

import (
	"Interview_Questions/Go-Zero/go-zero-courseware/user/rpc/userclient"
	"context"
	"google.golang.org/grpc/status"

	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/svc"
	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.UserInfoRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.UserInfoResponse{
		Id:        info.Id,
		UserName:  info.UserName,
		LoginName: info.LoginName,
		Sex:       info.Sex,
	}, nil
}
