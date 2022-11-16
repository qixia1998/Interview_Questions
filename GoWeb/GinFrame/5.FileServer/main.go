package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/file", fileServer)
	r.Run(":9090")
}

func fileServer(c *gin.Context) {
	path := "D:/go/src/Interview_Questions"
	fileName := path + c.Query("name")
	c.File(fileName)
}
