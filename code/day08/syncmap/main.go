package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.Map 是一个开箱即用的并发安全的map

var m = sync.Map{} // map 是结构体

func main() {
	wg := sync.WaitGroup{} // 空结构体初始化
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)         // 必须使用sync.Map内置的Store方法设置键值对
			value, _ := m.Load(key) // 必须使用sync.Map提供的Load方法根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
