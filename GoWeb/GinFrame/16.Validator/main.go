package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"net/http"
)

type ValUser struct {
	Name    string       `validate:"required" json:"name"`         // 非空
	Age     uint8        `validate:"gte=0, lte=130" json:"age"`    // 0<=Age<=130
	Email   string       `validate:"required, email" json:"email"` // 非空,email格式
	Address []ValAddress `validate:"dive" json:"address"`          // dive关键字代表 进入到嵌套结构体进行判断
}

type ValAddress struct {
	Province string `json:"province"`                         // 非空
	City     string `json:"city"`                             // 非空
	Phone    string `validate:"numeric, len=11" json:"phone"` // 数字类型，长度为11
}

var validate *validator.Validate

func init() {
	validate = validator.New() // 初始化（赋值）
}

func main() {
	r := gin.Default()
	var user ValUser
	r.POST("/validate", func(c *gin.Context) {
		testData(c)
		err := c.Bind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, "参数错误，绑定失败！")
			return
		}
		//执行参数的校验
		if validateUser(user) {
			c.JSON(http.StatusOK, "数据校验成功! ")
			return
		}
		c.JSON(http.StatusBadRequest, "数据校验失败! ")
	})
}

func testData(c *gin.Context) {
	address := ValAddress{
		Province: "广东省",
		City:     "深圳市",
		Phone:    "00000000000",
	}
	user := ValUser{
		Name:    "qixia",
		Age:     24,
		Email:   "qixia@gmail.com",
		Address: []ValAddress{address, address},
	}
	c.JSON(http.StatusOK, user)
}

func validateUser(u ValUser) bool {
	err := validate.Struct(u)
	if err != nil {
		// 断言为：validator.ValidationErrors, 类型为: []FieldError
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("错误的字段: ", e.Field())
			fmt.Println("错误的值: ", e.Value())
			fmt.Println("错误的tag ", e.Tag())
		}
		return false
	}
	return true
}
