package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

var sessionName string
var sessionValue string

type MyOption struct {
	sessions.Options
}

func main() {
	r := gin.Default()
	// 路由上加入session中间件
	store := cookie.NewStore([]byte("session_secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/session", func(c *gin.Context) {
		name := c.Query("name")
		if len(name) <= 0 {
			c.JSON(http.StatusBadRequest, "数据错误")
			return
		}
		sessionName = "session_" + name
		sessionValue = "session_value_" + name
		session := sessions.Default(c) // 获取的session
		sessionData := session.Get(sessionName)
		if sessionData != sessionValue {
			// 保存session
			session.Set(sessionName, sessionValue)
			o := MyOption{}
			o.Path = "/"
			o.MaxAge = 10 // 有效期 s
			session.Options(o.Options)
			session.Save() // 保存session
			c.JSON(http.StatusOK, "首次访问，session已经保存")
			return
		}
		c.JSON(http.StatusOK, "访问成功,您的session是: "+sessionData.(string))
	})
	r.Run(":9090")
}
