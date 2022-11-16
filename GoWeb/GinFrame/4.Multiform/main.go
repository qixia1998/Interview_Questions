package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// json格式输出
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello,Gin</b>",
		})
	})

	//原样输出html(html渲染)
	r.GET("/someHTML", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello,Gin</b>",
		})
	})

	//输出xml形式(XML渲染)
	r.GET("/someXML", func(c *gin.Context) {
		type Message struct {
			Name string
			Msg  string
			Age  int
		}
		info := Message{}
		info.Name = "qixia"
		info.Age = 24
		info.Msg = "Hello"
		c.XML(http.StatusOK, info)
	})

	//返回yaml形式(YAML渲染)
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "Gin框架多渲染形式",
			"status":  200,
		})
	})

	//开启服务
	r.Run(":9090")
}
