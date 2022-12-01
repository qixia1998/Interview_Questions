package main

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

var cookieName string
var cookieValue string

func main() {
	r := gin.Default()
	// Cookie中间件
	r.Use(CookieAuth())
	r.GET("/cookie", func(c *gin.Context) {
		name := c.Query("name")
		if len(name) <= 0 {
			c.JSON(http.StatusBadRequest, "数据错误")
			return
		}
		cookieName = "cookie_" + name
		cookieValue = hex.EncodeToString([]byte(cookieName + "value")) // cookie的值
		val, _ := c.Cookie(cookieName)
		if val == "" {
			c.String(http.StatusOK, "Cookie:%s已经下发, 下次登录有效", cookieName)
			return
		}
		c.String(http.StatusOK, "验证成功, cookie值为: %s", val)
	})
	r.Run(":9090")
}

func CookieAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		val, _ := c.Cookie(cookieName)
		if val == "" {
			c.SetCookie(cookieName, cookieValue, 3600, "/", "localhost", true, true)
		}
	}
}
