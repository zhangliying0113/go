package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个主 goroutine 去执行
func main() {
	for i := 0; i < 1000; i++ {
		go hello(i)
		// go func(i int) {
		// 	fmt.Println(i)
		// }(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	// main 函数结束了 由 main 函数启动的 goroutine也都结束了
}

// i 的值输出不是有序的，因为每次循环都会起一个 goroutine， goroutine 并行输出内容
