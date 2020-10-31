package main

import (
	"fmt"
	"strings"
)

// 闭包是一个函数，这个函数包含了他外部作用域的一个变量

// 底层原理：
// 函数可以作为返回值
// 函数内部查找变量的顺序，先在自己内部查找，找不到往外层找

func f1(f func()) {
	fmt.Println("this is f1")
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 要求 f1(f2)

func f3(fr func(int, int), x, y int) (fc func()) {
	fc = func() {
		fr(x, y)
	}
	return fc
}

// 闭包实例2
func adder1() func(int) int {
	var x int = 100
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包实例3
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 闭包实例4
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	fmt.Printf("%T\n", f2) // func(int, int)
	ret := f3(f2, 100, 200)
	fmt.Printf("%T\n", ret) // func()
	f1(ret)                 // this is f1

	ret2 := adder2(100)
	ret3 := ret2(200)
	fmt.Println(ret3) // 300

	ret4 := adder1()
	ret5 := ret4(200)
	fmt.Println(ret5) // 300

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")

	fmt.Println(jpgFunc("test"))    // test.jpg
	fmt.Println(txtFunc("xxx.txt")) // xxx.txt

	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) // 11 9
	fmt.Println(f1(3), f2(4)) // 12 8
	fmt.Println(f1(5), f2(6)) // 13 7

}
