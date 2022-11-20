package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		// 多文件上传
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusOK, "上传文件错误")
		}
		files := form.File["file_key"] // 上传的所有文件
		dst := "GoWeb/GinFrame/"
		// 遍历文件
		for _, file := range files {
			c.SaveUploadedFile(file, dst+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d 个文件上传完成！", len(files)))
	})
	r.Run(":9090")
}
