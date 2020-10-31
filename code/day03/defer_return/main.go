package main

import "fmt"

// go 语言中函数的 return 不是原子操作，在底层是分为两步执行
// 第一步：返回值赋值
// defer
// 第二部： 真正的 RET 返回

// 外城函数返回值改变的情况：
//1. 内部函数引用的是外部函数的参数并且值发生了改变
//2. 外部函数将值得地址传给内部函数并在内部改变了地址上的值

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 5 修改 x 前已经完成赋值
}

func f2() (x int) {
	defer func() {
		x++ // 修改 x 值时，因为当前函数没有 x，会直接修改 f2 中的 x，即 修改返回值
	}()
	return 5 // 6 返回值为 x，此时 x 值已变为6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 5 函数返回值为 y， 所以在返回值赋值是 y = x = 5，之后对 x 改变后，y 值不变
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 5 defer 函数本身有 x， defer 改变的是内部函数的 x， return x 结果为 x = 5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5 // 5 defer 函数中传入的 x 被改变，但是作用域是 defer 函数，传的是值，不会改变外层 x
}

func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 // 6 defer 函数传入的是外层函数的地址，x 地址上的值被改变，外层变量也会发生改变
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
}
