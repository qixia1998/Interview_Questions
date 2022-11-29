package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Response struct {
	code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func main() {
	r := gin.Default()
	r.GET("/login", login)
	r.POST("/register", register)
	r.Run(":9090")
}

func register(c *gin.Context) {
	var user User
	err := c.Bind(&user)
	if err != nil {
		fmt.Println("绑定错误: ", err)
		c.JSON(http.StatusBadRequest, "数据错误！")
		return
	}
	res := Response{
		code: http.StatusOK,
		Msg:  "注册成功",
		Data: "Ok",
	}
	c.JSON(http.StatusOK, res)
}

func login(c *gin.Context) {
	userName := c.Query("name")
	pwd := c.Query("pwd")
	fmt.Println(userName, pwd)
	res := Response{}
	res.code = http.StatusOK
	res.Msg = "登录成功"
	res.Data = "Ok"
	c.JSON(http.StatusOK, res)
}
