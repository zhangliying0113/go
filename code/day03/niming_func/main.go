package main

import "fmt"

// 匿名函数
// 函数内部没有办法声明带名字的函数

func main() {
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	// 如果只是调用一次的函数，还可以简写成立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("hello go")
	}(100, 200)

}
