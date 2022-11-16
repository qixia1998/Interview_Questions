package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/file", fileServer)
	r.Run(":9090")
}

func fileServer(c *gin.Context) {
	path := "D:/go/src/Interview_Questions/GoWeb/GinFrame/5.FileServer/"
	fileName := path + c.Query("name")
	c.File(fileName)
}
