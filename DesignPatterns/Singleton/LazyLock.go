package Singleton

import "sync"

type singleton struct {
}

var ins *singleton
var mu sync.Mutex

func GetIns() *singleton {
	if ins == nil {
		mu.Lock()
		if ins == nil {
			ins = &singleton{}
		}
		mu.Unlock()
	}
	return ins
}
