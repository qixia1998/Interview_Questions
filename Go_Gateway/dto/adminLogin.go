package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/qixia1998/Interview_Questions/Go_Gateway/public"
)

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"required,is_valid_username"` // 管理员用户名
	PassWord string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required" `                 // 密码
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""` // token
}
