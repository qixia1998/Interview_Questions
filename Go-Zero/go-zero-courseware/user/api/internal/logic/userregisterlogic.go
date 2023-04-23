package logic

import (
	"Interview_Questions/Go-Zero/go-zero-courseware/user/rpc/userclient"
	"context"
	"google.golang.org/grpc/status"

	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/svc"
	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	_, err = l.svcCtx.UserRpc.Register(l.ctx, &userclient.RegisterRequest{
		LoginName: req.LoginName,
		UserName:  req.UserName,
		PassWord:  req.PassWord,
		Sex:       req.Sex,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.RegisterResponse{}, nil
}
