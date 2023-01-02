package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pjebs/restgate"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(authMiddleware())
	r.GET("/auth1", func(c *gin.Context) {
		resData := struct {
			Code int         `json:"code"`
			Msg  string      `json:"msg"`
			Data interface{} `json:"data"`
		}{http.StatusOK, "验证通过", "OK"}
		c.JSON(http.StatusOK, resData)
	})
	r.Run(":9090")
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		gate := restgate.New("X-Auth-Key", "X-Auth-Secret",
			restgate.Static,
			restgate.Config{
				Key:                []string{"admin", "gin"},
				Secret:             []string{"adminpw", "gin_ok"},
				HTTPSProtectionOff: true,
			})
		nextCalled := false
		nextAdapter := func(http.ResponseWriter, *http.Request) {
			nextCalled = true
			c.Next()
		}
		gate.ServeHTTP(c.Writer, c.Request, nextAdapter)
		if nextCalled == false {
			c.AbortWithStatus(401)
		}
	}
}
