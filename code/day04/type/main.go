package main

import "fmt"

// 自定义类型和类型别名

// type 后面跟的是类型
type myInt int    // 自定义类型
type youInt = int // 为 int 类型设置别名

func main() {
	var n myInt
	n = 100
	fmt.Println(n)               // 100
	fmt.Printf("%T  %d\n", n, n) // main.myInt  100

	var m youInt
	m = 100
	fmt.Println(m)               // 100
	fmt.Printf("%T  %d\n", m, m) // int  100

	var c rune
	c = '中'
	fmt.Println(c)                      // 100
	fmt.Printf("%T  %d  %c\n", c, c, c) // int32  20013  中

}
