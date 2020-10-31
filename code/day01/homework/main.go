package main

import (
	"fmt"
	"unicode"
)

func main() {
	// 1. 常量声明练习，特别是iota定义数量级时（1024）

	// iota 只能在常量表达式中使用
	// 每次 const 出现时，都会让 iota 初始化为0，每换一行 iota 自增长
	const (
		a = 1
		b = iota
	)

	const (
		x = iota // 0
		y        // 1
		z        // 2
	)

	const (
		u = iota // 0
		v        // 1
		_
		_
		w //4
	)

	// iota 用于位掩码表达式
	const (
		IgEggsAllergen = 1 << iota // 1 << 0 which is 00000001
		IgChocolate                // 1 << 1 which is 00000010
		IgNuts                     // 1 << 2 which is 00000100
		IgStrawberries             // 1 << 3 which is 00001000
		IgShellfish                // 1 << 4 which is 00010000
	)

	// 定义数量级
	const (
		_          = iota             // ignore first value by assigning to blank identifier
		KBByteSize = 1 << (10 * iota) // 1 << (10*1)
		MB                            // 1 << (10*2)
		GB                            // 1 << (10*3)
		TB                            // 1 << (10*4)
		PB                            // 1 << (10*5)
		EB                            // 1 << (10*6)
		ZB                            // 1 << (10*7)
		YB                            // 1 << (10*8)
	)

	// 2. 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用`fmt.Printf()`搭配`%T`分别打印出上述变量的值和类型。
	var age = 20
	var score = 80.7
	var isOk = false
	name := "神州"

	fmt.Printf("%d %T\n", age, age)
	fmt.Printf("%f %T\n", score, score)
	fmt.Printf("%t %T\n", isOk, isOk)
	fmt.Printf("%s %T\n", name, name)

	// 3. 编写代码统计出字符串`"hello沙河小王子"`中汉字的数量。(自己查资料)
	hanString := `"hello沙河小王子"`
	count := 0
	for _, v := range hanString {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println(count)

	// 4. 九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
