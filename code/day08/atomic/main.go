package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x) // 10000

	// 比较并交换
	ok := atomic.CompareAndSwapInt64(&x, 10000, 300)
	fmt.Println(ok, x) // true 300

}
