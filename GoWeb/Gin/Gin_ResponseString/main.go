package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/response/", ResponseStringHandler)
	fmt.Println("http://127.0.0.1:8000/response/")
	r.Run(":8000")
}

func ResponseStringHandler(c *gin.Context) {
	c.String(http.StatusOK, "返回字符串")
}
