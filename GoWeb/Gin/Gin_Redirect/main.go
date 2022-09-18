package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1")
	})
	fmt.Println("http://127.0.0.1:8000")
	r.Run(":8000")
}
