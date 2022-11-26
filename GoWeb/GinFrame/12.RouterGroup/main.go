package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResGroup struct {
	Data string
	Path string
}

func main() {
	router := gin.Default()
	// 路由分组1
	v1 := router.Group("/v1") // 路由分组（1级路径）
	{
		r := v1.Group("/user") // 路由分组（2级路径）
		r.GET("/login", login) // 响应请求:/v1/user/login
		// 路由分组 （3级路径）
		r2 := r.Group("showInfo")     // /v1/user/showInfo
		r2.GET("/abstract", abstract) // /v1/user/showInfo/abstract
		r2.GET("/detail", detail)     // /v1/user/showInfo/detail
	}
	// 路由分组2
	v2 := router.Group("/v2")
	{
		v2.GET("/other", other) //响应请求: /v2/other
	}

	router.Run(":9090")

}

func other(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{"other", c.Request.URL.Path})
}

func detail(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{"detail", c.Request.URL.Path})
}

func abstract(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{"abstract", c.Request.URL.Path})
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{"login", c.Request.URL.Path})
}
