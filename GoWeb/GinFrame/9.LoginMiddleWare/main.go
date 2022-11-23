package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 路由使用 gin.BasicAuth() 中间件
	r.Use(AuthMiddleware())
	r.GET("/login", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		c.JSON(http.StatusOK, "登录成功!"+"欢迎您: "+user)
	})
	// 监听并在 0.0.0.0:9090 上启动服务
	r.Run(":9090")
}

func AuthMiddleware() gin.HandlerFunc {
	// 初始化用户
	accounts := gin.Accounts{ // gin.Accounts 是 map[string]string 类型
		"admin":  "admin",
		"system": "system",
	}
	// 动态添加用户
	accounts["go"] = "123456789"
	accounts["gin"] = "gin123"
	// 将用户添加到登录中间件中
	auth := gin.BasicAuth(accounts)
	return auth
}
