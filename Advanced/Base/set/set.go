package main

import "fmt"

// 切片元素的去重
// 使用map的key

func main() {
	s1 := []string{"abc", "def", "abc", "ok", "ok"}

	m := make(map[string]struct{})
	for _, i := range s1 {
		m[i] = struct{}{}
	}
	s2 := make([]string, 0)
	for k := range m {
		s2 = append(s2, k)
	}
	fmt.Println(s1)
	fmt.Println(s2)
}
