package main

import "fmt"

func main() {
	m := make(map[float64]int)
	m[2.4] = 2
	fmt.Printf("k: %v, v: %d\n", 2.4000000000000000000000001, m[2.4000000000000000000000001])
	fmt.Println(m[2.4000000000000000000000001] == m[2.4])
}
