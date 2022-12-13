package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"net/http"
)

type HttpRes struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func main() {
	r := gin.Default()
	r.GET("/https_test", func(c *gin.Context) {
		fmt.Println(c.Request.Host)
		c.JSON(http.StatusOK, HttpRes{
			Code:   http.StatusOK,
			Result: "测试成功",
		})
	})
	path := "xxx"                                   // 证书路径
	r.RunTLS(":9090", path+"xx.crt", path+"xx.key") // 开启HTTPS服务
}

func httpsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		secureMiddle := secure.New(secure.Options{
			SSLRedirect: true, // 只允许https请求
			//SSLHost:     "",   // http到https的重定向
			STSSeconds:           1536000, // Strict-Transport-Security header的时效：1年
			STSIncludeSubdomains: true,
			STSPreload:           true, // STS Preload(预加载)
			FrameDeny:            true,
			ContentTypeNosniff:   true, // 禁用浏览器的类型猜测行为，防止基于 MIME 类型混淆的攻击
			BrowserXssFilter:     true, // 启用XSS保护，并在检查到XSS攻击时，停止渲染页面
			//IsDevelopment: true, // 开发模式
		})
		err := secureMiddle.Process(context.Writer, context.Request)
		// 如果不安全，终止
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "数据不安全")
			return
		}
		// 如果是重定向，终止

		if status := context.Writer.Status(); status > 300 && status < 399 {
			context.Abort()
			return
		}
		context.Next()
	}
}
