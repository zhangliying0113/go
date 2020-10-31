package main

import "fmt"

// defer 多用于函数结束之前释放资源（文件句柄、数据库连接、socket 连接）
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("111") // defer 把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("222") // 一个函数中可以有多个defer语句
	defer fmt.Println("333") // 多个defer语言按照先进后出（后进先出）的顺序延迟执行
	fmt.Println("end")
}

func main() {
	deferDemo()
}

// run:
// start
// end
// 333
// 222
// 111
