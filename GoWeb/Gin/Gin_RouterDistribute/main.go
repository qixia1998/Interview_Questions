package main

import (
	"Interview_Questions/GoWeb/Gin/Gin_RouterDistribute/routers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.LoadUsers(r)
	routers.LoadBooks(r)
	fmt.Println("用户路由：http://127.0.0.1:8000/user")
	fmt.Println("书籍路由：http://127.0.0.1:8000/book")
	r.Run(":8000")
}
