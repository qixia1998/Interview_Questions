package main

import "fmt"

// 只声明
//var m1 map[string]string

// 声明又初始化
//var m2 = map[string]string{"a": "b"}

func main() {
	m := make(map[string]int)
	keys := make([]string, 0)
	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("key_%d", i)
		keys = append(keys, key)
		m[key] = i
	}
	fmt.Println(m)
	// range 遍历 keys
	//for k := range m {
	//	fmt.Printf("[key=%s]\n", k)
	//}

	fmt.Println("无序遍历")
	// range 遍历
	for k, v := range m {
		fmt.Printf("[%s=%d]\n", k, v)
	}
	// 有序key遍历
	fmt.Println("有序遍历")
	for _, k := range keys {
		fmt.Printf("[%s=%d]\n", k, m[k])
	}
}
