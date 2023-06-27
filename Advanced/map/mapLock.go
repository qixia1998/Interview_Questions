package main

import (
	"fmt"
	"sync"
	"time"
)

// 为了解决map线程不安全 ，我们自己加锁

type concurrentMap struct {
	mp map[int]int
	sync.RWMutex
}

// 通过set 方法做原有map的赋值 m[key] =v
func (c *concurrentMap) Set(key, value int) {
	// 加写锁
	c.Lock()
	c.mp[key] = value
	c.Unlock()

}

// 通过get 方法做原有map的读取值操作 v:= m[key]
func (c *concurrentMap) Get(key int) int {
	//先获取读锁
	c.RLock()
	res := c.mp[key]
	c.RUnlock()
	return res
}

func main() {
	c := concurrentMap{
		mp: make(map[int]int),
	}
	// 一个线程循环写map
	go func() {
		for i := 0; i < 10000; i++ {
			c.Set(i, i)
		}
	}()
	// 一个线程循环读map
	go func() {
		for i := 0; i < 10000; i++ {
			res := c.Get(i)
			fmt.Printf("[cmap.get][%d=%d]\n", i, res)
		}
	}()
	time.Sleep(1 * time.Hour)

}
