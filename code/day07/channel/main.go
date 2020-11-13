package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int // 需要制定通道中元素的类型
var wg sync.WaitGroup

// 不带缓冲区通道的初始化
func noBufChannel() {
	b = make(chan int)
	fmt.Println(b) // 0xc000018120 开辟内存空间
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b // 如果未在 b 通道放值直接取值会报错：fatal error: all goroutines are asleep - deadlock!
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()

	b <- 10
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
}

// 带缓冲区通道的初始化
func bufChannel() {
	b = make(chan int, 10)
	b <- 10
	fmt.Println("10发送到通道b中了...")
	b <- 20
	fmt.Println("20发送到通道b中了...")
	x := <-b
	fmt.Println("从通道b中取到了", x)
	close(b)
}

func main() {
	fmt.Println(b) // <nil>
	// noBufChannel()
	bufChannel() // 两个值都能发送进去
	// 10发送到通道b中了...
	// 20发送到通道b中了...
	// 从通道b中取到了 10
}
