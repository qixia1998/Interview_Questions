package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/response/", ResponseJsonHandler)
	fmt.Println("http://127.0.0.1:8000/response/")
	r.Run(":8000")
}

func ResponseJsonHandler(c *gin.Context) {
	type Data struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}
	d := Data{
		Msg:  "Json数据",
		Code: 1001,
	}
	c.JSON(http.StatusOK, d)

	//// 也可以直接使用 gin.H返回 json数据
	//c.JSON(http.StatusOK, gin.H{
	// "msg": "success",
	//})
}
