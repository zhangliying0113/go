package main

import "fmt"

// 函数：一段代码的封装
func f1() {
	fmt.Println("hello go")
}

func f2(name string) {
	fmt.Println("hello", name)
}

func f3(x int, y int) int {
	sum := x + y
	return sum
}

func f4(x, y int) int {
	return x + y
}

// 可变参数
func f5(title string, y ...int) int {
	fmt.Println(y) // y 是一个 int 类型切片
	return 1
}

// 命名返回值
func f6(x, y int) (sum int) {
	sum = x + y // 如果使用命名的返回值，那么在函数中可以直接使用返回值变量
	return      // 如果使用命名的返回值,return后面可以省略返回值变量
}

// go 语言支持多个返回值
func f7(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f2("lixiang")
	fmt.Println(f3(100, 200))
	f5("liixang", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum, sub := f7(200, 100)
	fmt.Println(sum, sub)

	// 在一个命名的函数中不能够再声明命名函数
	// func f8(){

	// }
}
