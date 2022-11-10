package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 一般重定向: 重定向到外部网络
	r.GET("/redirect1", func(c *gin.Context) {
		// 重定向状态码:
		url := "http://www.baidu.com"
		c.Redirect(http.StatusMovedPermanently, url)
	})
	r.Run(":9090")
}
