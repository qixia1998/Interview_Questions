package logic

import (
	"Interview_Questions/Go-Zero/go-zero-courseware/user/rpc/model"
	"context"
	"google.golang.org/grpc/status"

	"Interview_Questions/Go-Zero/go-zero-courseware/user/rpc/internal/svc"
	"Interview_Questions/Go-Zero/go-zero-courseware/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	_, err := l.svcCtx.UserModel.FindOneByLoginName(l.ctx, in.LoginName)
	if err == nil {
		return nil, status.Error(5000, "登录名已存在")
	}

	if err != model.ErrNotFound {
		return nil, status.Error(500, err.Error())
	}
	newUser := model.User{
		LoginName: in.LoginName,
		Username:  in.UserName,
		Sex:       in.Sex,
		Password:  in.PassWord,
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &newUser)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &user.RegisterResponse{}, nil
}
