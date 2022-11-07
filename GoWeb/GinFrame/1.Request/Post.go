package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()       // 路由引擎
	r.POST("/post", postMsg) // post方法
	r.Run(":9090")           // 本地IP地址，9090端口
}

func postMsg(c *gin.Context) {
	// name := c.Query("name") // 获取URL中的数据
	name := c.DefaultPostForm("name", "qixia")
	fmt.Println(name)
	from, b := c.GetPostForm("name")
	fmt.Println(from, b)
	c.JSON(http.StatusOK, "欢迎您:"+name)
}
