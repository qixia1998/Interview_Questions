package main

import (
	"log"
	"sync"
	"time"
)

// HcMutex 是一个互斥锁
var HcMutex sync.Mutex

func runMutex(id int) {
	log.Printf("[任务id :%d][尝试获取锁]", id)
	HcMutex.Lock()
	log.Printf("[任务id :%d][获取到了锁]", id)
	time.Sleep(20 * time.Second)
	HcMutex.Unlock()
	log.Printf("[任务id :%d][执行完成 释放锁]", id)
}

func runHcLock() {
	go runMutex(1)
	go runMutex(2)
	go runMutex(3)
}

func main() {
	runHcLock()

	time.Sleep(6 * time.Minute)
}
