package main

import "fmt"

// go 语言函数外的语句必须使用关键字开
// go 语言中推荐使用小驼峰命名
var studentName string

// 批量声明变量
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "双引号"
	age = 16
	isOk = false

	// go 语言中非全局变量声明后必须使用，不使用就编译不过去
	// var hy string = "jinzuanshi"
	fmt.Println(name)
	fmt.Println()                 // Println:打印完指定内容后会在后面加一个换行符，无参数代表打印一个空行
	fmt.Printf("name:%s\n", name) // %s:string 类型占位符

	// 声明变量的同事赋值
	var s1 string = "go"
	println(s1)

	// 类型推到（根据值判断该变量是什么类型）
	var s2 = 20
	println(s2)

	// 简短变量声明，只能在函数里面用
	s3 := "语言"
	println(s3)
	// 同一个作用域（{}）中不能重复声明同名的变量
	// 匿名变量是一个特殊的变量：_
	// 当有些数据必须用变量接收但是又不使用它时，就可以用（哑元变量） _ 来接收这个值。
}
