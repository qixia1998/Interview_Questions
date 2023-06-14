package main

import "fmt"

// 异或的方式交换两个变量

func main() {
	var a = 100
	var b = 200
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Printf("a=%d, b=%d", a, b)
}
