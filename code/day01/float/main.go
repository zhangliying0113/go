package main

import "fmt"

func main() {
	// go 语言中默认小数都是 float64 位
	f1 := 3.141
	fmt.Printf("%T\n", f1)

	f2 := float32(3.141)
	fmt.Printf("%T\n", f2)

	// f1 = f2    float32 类型的值不能直接赋值给 float64 类型的变量
}
