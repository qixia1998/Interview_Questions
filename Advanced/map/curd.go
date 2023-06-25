package main

import "fmt"

// 仅声明
var m1 map[string]string

// 声明并初始化
var m2 = map[string]string{"a": "b"}

func main() {
	m := make(map[string]string)
	// 增加
	m["tool"] = "goland"
	m["lang"] = "golang"
	// 删除
	delete(m, "tool")
	fmt.Println(m)
	// 改
	m["lang"] = "python"
	// 查
	// 单变量形式
	lang := m["lang"]
	fmt.Println(lang)
	// 双变量形式
	lang1, exists := m["lang"]
	if exists {
		fmt.Printf("[lang存在 值：%v]\n", lang1)
	} else {
		fmt.Println("lang不存在")
		m["lang"] = "java"
	}

}
