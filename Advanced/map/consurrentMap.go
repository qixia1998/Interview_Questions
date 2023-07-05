package main

import (
	"fmt"
	"github.com/orcaman/concurrent-map"
	"log"
	"time"
)

func main() {
	// Create a new map.
	m := cmap.New()

	// 循环写map
	go func() {

		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			m.Set(key, i)
		}

	}()
	// 循环读map
	go func() {

		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			v, exists := m.Get(key)
			if exists {
				log.Printf("[%s=%v]", key, v)
			}
		}

	}()
	// 循环写map
	go func() {

		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			m.Set(key, i)
		}

	}()
	// 循环写map
	go func() {

		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			m.Set(key, i)
		}

	}()
	// 循环写map
	go func() {

		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			m.Set(key, i)
		}

	}()

	time.Sleep(1 * time.Hour)

}
