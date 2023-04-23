package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"Interview_Questions/Go-Zero/go-zero-courseware/user/rpc/internal/svc"
	"Interview_Questions/Go-Zero/go-zero-courseware/user/rpc/model"
	"Interview_Questions/Go-Zero/go-zero-courseware/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err == model.ErrNotFound {
		return nil, status.Error(5000, "用户不存在")
	}
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &user.UserInfoResponse{
		Id:        userInfo.Id,
		UserName:  userInfo.Username,
		LoginName: userInfo.LoginName,
		Sex:       userInfo.Sex,
	}, nil
}