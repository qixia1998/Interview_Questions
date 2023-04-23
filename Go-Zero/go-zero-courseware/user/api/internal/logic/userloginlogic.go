package logic

import (
	"Interview_Questions/Go-Zero/go-zero-courseware/user/rpc/userclient"
	"context"

	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/svc"
	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	login, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginRequest{
		LoginName: req.LoginName,
		Password:  req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Id:    login.Id,
		Token: login.Token,
	}, nil
}
