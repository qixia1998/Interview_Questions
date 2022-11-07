package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()    // 路由引擎
	r.GET("/get", getMsg) //get方法
	//r.Run("127.0.0.1:9090")
	r.Run(":9090") // 默认本机IP地址，8080端口
}

func getMsg(c *gin.Context) {
	name := c.Query("name")
	// 返回字符串
	//c.String(http.StatusOK, "欢迎您:%s", name)
	// 返回json数据
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "返回信息",
		"data": "欢迎您" + name,
	})
}
