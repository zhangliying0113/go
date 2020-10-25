package main

import "fmt"

// 函数定义 入参 出参
func sum(x int, y int) (ret int) {
	return x + y
}

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数有返回值
func f3() int {
	ret := 3
	return ret
}

// 返回值可以命名也可以不命名
// 命名的返回值相当于在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return // 使用命名返回值可以 return 后省略
}

// 多个返回值
func f5() (int, string) {
	return 1, "shahe"
}

// 参数的类型简写
// 当参数中连续多个参数类型一致时，我们可以将非最后一个参数的类型省略
func f6(x, y, z int, m, n string, f, t bool) int {
	return x + y
}

// 可变长参数
// 可变长参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y 的类型是切片 []int
}

// go 语言函数没有默认参数这个概念

func main() {
	f7("shijiazhuang", 1, 2, 3, 4, 5, 6, 7) // 直接传入切片的值
}
