package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Remark   string `json:"remark"`
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var login Login
		err := c.Bind(&login) // 执行绑定
		fmt.Println("绑定的数据: ", login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "绑定失败, 参数错误",
				"data": err.Error(),
			})
			return
		}
		if login.UserName == "user" && login.Password == "123456" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "登录成功",
				"data": "OK",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "登录失败",
			"data": "error",
		})
		return
	})
	r.Run(":9090")
}
