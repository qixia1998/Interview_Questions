package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/GetOtherData", func(c *gin.Context) {
		url := "http://www.baidu.com"
		//url := "https://img2.baidu.com/it/u=645336210,1332246460&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=1082"
		response, err := http.Get(url)
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable) // 应答client
			return
		}
		body := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		// 数据写入响应体
		c.DataFromReader(http.StatusOK, contentLength, contentType, body, nil)
	})
	r.Run(":9090")
}
