package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorldHandler(c *gin.Context) {
	// gin.Context,封装了request和response
	c.String(http.StatusOK, "Hello World!")
}

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// 基本路由 /hello/
	r.GET("/hello", HelloWorldHandler)
	// 3.监听端口，默认在8080
	fmt.Println("运行地址: http://127.0.0.1:8000")
	r.Run(":8000")
}
