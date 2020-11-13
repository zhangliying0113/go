package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS

var wg sync.WaitGroup

func f() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func f1() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.NumCPU()) // 12 返回当前系统的 CPU 核数量
	wg.Add(2)
	go f()
	go f1()
	wg.Wait()
}
