package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

func main() {

	m := sync.Map{}
	// 新增
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		m.Store(key, i)
	}
	// 删除
	m.Delete("key_8")

	// 修改m.Store
	m.Store("key_9", 999)

	// 查询
	res, loaded := m.Load("key_09")
	if loaded {
		//  类型断言 res.(int)
		log.Printf("[key_09存在 :%v 数字类型:%d]", res, res.(int))
	}

	// 遍历 return false 停止
	m.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(int)
		if strings.HasSuffix(k, "3") {
			log.Printf("不想要3")
			//return true
			return false
		} else {
			log.Printf("[sync.map.Range][遍历][key:=%s][v:=%d]", k, v)

			return true
		}

	})
	// LoadAndDelete 先获取值再删掉
	s1, loaded := m.LoadAndDelete("key_7")
	log.Printf("key_7 LoadAndDelete :%v", s1)
	s2, loaded := m.Load("key_7")
	log.Printf("key_7 LoadAndDelete:%v", s2)

	actual, loaded := m.LoadOrStore("key_8", 158)
	if loaded {
		log.Printf("key_8原来的值是:%v", actual)
	} else {
		log.Printf("key_8原来没有，实际是:%v", actual)
	}

	actual, loaded = m.LoadOrStore("key_1", 158)
	if loaded {
		log.Printf("key_1原来的值是:%v", actual)
	} else {
		log.Printf("key_1原来没有，实际是:%v", actual)
	}
}
