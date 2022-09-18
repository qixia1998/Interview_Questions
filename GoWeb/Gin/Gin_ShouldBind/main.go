package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `from:"password" json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var login Login
	if err := c.ShouldBind(&login); err != nil {
		// 如果数据校验不通过直接返回
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.String(http.StatusOK, fmt.Sprintf("姓名：%s -- 密码：%s", login.Username, login.Password))
}

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// 基本路由 /login/
	r.POST("/login/", LoginHandler)
	// 3.监听端口，默认在8080
	fmt.Println("运行地址：http://127.0.0.1:8000/login/")
	r.Run(":8000")
}
