package main

// panic可以在任何地方引发，但recover只有在defer调用的函数中有效
// recover()必须搭配defer使用。
// defer一定要在可能引发panic的语句之前定义。

import "fmt"

func funcA() {
	fmt.Println("a")
}

func funcB() {
	// 刚刚打开数据库连接
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接。。。")
	}()
	panic("出现了严重的错误！！！") // 程序崩溃退出
	fmt.Println("b")     // 绿色下划线是因为程序在 panic 之后，不会被执行到，是无效代码
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
