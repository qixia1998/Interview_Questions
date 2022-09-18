package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadBooks(e *gin.Engine) {
	e.GET("/book", GetBookHandler)
}

func GetBookHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Book Router",
	})
}
