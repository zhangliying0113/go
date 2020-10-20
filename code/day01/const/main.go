package main

import "fmt"

// 常量声明
const pi = 3.1415926

// 批量声明常量
const (
	statusOk = 200
	notFount = 404
	name     = "go"
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行值一样
const (
	n1 = 100
	n2
	n3
)

// iota
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

const (
	b1 = iota // 0
	b2 = iota // 1
	_  = iota // 2
	b3 = iota // 3
)

const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4        // 3
)

const (
	d1, d2 = iota + 1, iota + 2 // d1:1 d2:2

	// 中间空多行代表换一行
	d3, d4 = iota + 1, iota + 2 // d3:2 d4:3
)

// 定义数量级
const (
	_  = iota             // 0
	KB = 1 << (10 * iota) // 1的10次方*2
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	fmt.Println("---------------")

	var n = 100
	// 查看类型
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	var s = "Hello 沙河！"
	fmt.Printf("字符串：%s\n", s)
	fmt.Printf("字符串：%v\n", s)
	fmt.Printf("字符串：%#v\n", s) // # 将字符串双引号打印出来
}
