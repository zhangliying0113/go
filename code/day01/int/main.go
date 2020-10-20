package main

import "fmt"

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1) // %d:十进制
	fmt.Printf("%b\n", i1) // %b:二进制
	fmt.Printf("%o\n", i1) // %o:八进制
	fmt.Printf("%x\n", i1) // %x:十六进制

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	// 声明 int8 类型的变量
	i4 := int8(9) // 明确指定 int8 类型，否则就默认为 int 类型
	fmt.Printf("%T\n", i4)
}
