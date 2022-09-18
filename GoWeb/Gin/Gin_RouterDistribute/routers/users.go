package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadUsers(e *gin.Engine) {
	e.GET("/user", UserHandler)
}

func UserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User Router",
	})
}
