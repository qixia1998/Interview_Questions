package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("fileName")
		if err != nil {
			c.String(http.StatusOK, "文件上传错误！")
		}
		// 存储路径
		dst := ""
		c.SaveUploadedFile(file, dst+file.Filename) // 存储文件
		c.String(http.StatusOK, fmt.Sprintf("%s 上传完成", file.Filename))
	})
	r.Run(":9090")
}
