package main

import (
	"Interview_Questions/Security/utils"
	"fmt"
)

func main() {

	src := "hello"
	ret := utils.Base64Encoding(src)
	fmt.Println(ret)

	retByte, err := utils.Base64Decoding(ret)
	if err != nil {
		fmt.Println("Decoding error")
	}
	fmt.Println(retByte)
}
