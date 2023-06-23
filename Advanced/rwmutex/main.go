package main

import (
	"log"
	"sync"
	"time"
)

var rwMutex sync.RWMutex

func runReadLock(id int) {
	log.Printf("[读任务id :%d][进入读方法尝试获取读锁]", id)
	rwMutex.RLock()
	log.Printf("[读任务id :%d][获取到了读锁][开始执行 睡眠10秒]", id)
	time.Sleep(10 * time.Second)
	rwMutex.RUnlock()
	log.Printf("[读任务id :%d][完成读任务 释放读锁]", id)
}

func runWriteLock(id int) {
	log.Printf("[写任务id :%d][进入写方法尝试获取写锁]", id)
	rwMutex.Lock()
	log.Printf("[写任务id :%d][获取到了写锁][开始执行 睡眠10秒]", id)
	time.Sleep(10 * time.Second)
	rwMutex.Unlock()
	log.Printf("写任务id :%d][完成写任务 释放写锁]", id)
}

// 全是写任务
func allWriteWorks() {
	for i := 1; i < 6; i++ {
		go runWriteLock(i)
	}
}

// 全是读任务
func allReadWorks() {
	for i := 1; i < 6; i++ {
		go runReadLock(i)
	}
}

// 先启动写任务
func writeFirst() {
	go runWriteLock(1)
	time.Sleep(1 * time.Second)
	go runReadLock(1)
	go runReadLock(2)
	go runReadLock(3)
	go runReadLock(4)
	go runReadLock(5)
}

// 先启动读任务
func readFirst() {
	go runReadLock(1)
	go runReadLock(2)
	go runReadLock(3)
	go runReadLock(4)
	go runReadLock(5)

	time.Sleep(1 * time.Second)
	go runWriteLock(1)
}

func main() {
	log.Printf("执行读写锁效果的函数")
	// 1.同时更多个写锁任务，说明如果并非使用读写锁的写锁时，退化成了互斥锁
	// allWriteWorks()

	// 2.同时多个读锁任务，说明使用读写锁的读锁，可以同时施加多把读锁
	// allReadWorks()

	// 3.先启动写锁任务，后并发5个读锁任务，当有写锁存在时，读锁是施加不了的。写锁释放完，读锁可以施加多个
	// writeFirst()

	// 4.先并发5个读锁任务，后启动一个写锁任务，当有读锁时，阻塞写锁。
	readFirst()

	time.Sleep(1 * time.Hour)
}
