package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserDetailHandler(c *gin.Context) {
	//username := c.DefaultQuery("name", "xxx")
	username := c.Query("name")
	// gin.Context，封装了request和response
	c.String(http.StatusOK, fmt.Sprintf("姓名：%s", username))
}

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// 基本路由 /user?name=root
	r.GET("/user/", GetUserDetailHandler)
	// 3.监听端口，默认在8080
	fmt.Println("运行地址: http://127.0.0.1:8000")
	r.Run(":8000")
}
