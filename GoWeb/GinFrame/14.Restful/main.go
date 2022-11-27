package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

// 调用第三方接口的请求数据
type UserAPI struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type TempData struct {
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

// 客户端提交的数据
type ClientRequest struct {
	UserName string      `json:"user_name"`
	Password string      `json:"password"`
	Other    interface{} `json:"other"`
}

// 返回客户端的数据
type ClientResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func main() {
	//testAPi()
	r := gin.Default()
	r.POST("/getOtherAPI", getOtherAPI)
	r.Run(":9091")
}

func getOtherAPI(context *gin.Context) {
	var requestData ClientRequest
	var response ClientResponse
	err := context.Bind(&requestData)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Msg = "请求的参数错误"
		response.Data = err
		context.JSON(http.StatusBadRequest, response)
		return
	}
	// 请求第三方API接口数据
	url := "http://127.0.0.1:9090/login"
	user := UserAPI{requestData.UserName, requestData.Password}
	data, err := getRestfulAPI(url, user, "application/json")
	fmt.Println(data, err)
	var temp TempData
	json.Unmarshal(data, &temp)
	fmt.Println(temp.Msg, temp.Data)
	response.Code = http.StatusOK
	response.Msg = "请求数据成功"
	response.Data = temp
	context.JSON(http.StatusOK, response)
}

func testAPi() {
	url := "http://127.0.0.1:9090/login"
	user := UserAPI{"user", "123456"}
	data, err := getRestfulAPI(url, user, "application/json")
	fmt.Println(data, err)
	var temp TempData
	json.Unmarshal(data, &temp)
	fmt.Println(temp.Msg, temp.Data)
}

// 发送POST请求
// url:  请求地址
// data:  POST请求提交的数据
// contentType:  请求体格式, 如: application/json
// content: 请求返回的内容
func getRestfulAPI(url string, data interface{}, contentType string) ([]byte, error) {
	// 创建调用API接口的client
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("调用API接口出现了错误！")
		return nil, err
	}
	res, err := ioutil.ReadAll(resp.Body)
	return res, err
}
