package main

import (
	"Interview_Questions/Security/utils"
	"fmt"
)

func main() {

	src := "hello"
	//src := "127.0.0.1:8000/?name=hello"
	ret := utils.Base64Encoding(src)
	//ret := utils.Base64UrlEncoding(src)
	fmt.Println(ret)

	retByte, err := utils.Base64Decoding(ret)
	//retByte, err := utils.Base64UrlDecoding(ret)
	if err != nil {
		fmt.Println("Decoding error")
	}
	fmt.Println(retByte)
}
