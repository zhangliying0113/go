package main

import "fmt"

// Go语言中为了处理非ASCII码类型的字符 定义了新的rune类型
// rune: 复合字符类型,即 int32 类型。
// byte: uint8 类型，字节，一个英文字母代表一个字节。
func main() {
	s := "Hello沙河사샤"
	// len() 求的是 byte 字节的数量
	n := len(s) // n = 17,是指的17个字节，一个汉字或者其他复合字符占多个字节
	fmt.Println(n)

	// 遍历字节
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}

	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

	s2 := "白萝卜"
	s3 := []rune(s2) // 字符串强制转换一个 rune 切片
	s3[0] = '红'
	fmt.Println(string(s3)) // 把 rune 切片转换成字符串

	c1 := "红" // byte 类型
	c2 := '红' // rune 类型
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	c3 := "H"       // string
	c4 := byte('H') // byte(uint8)
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("%d\n", c4)

	// 类型转换
	n1 := 10 // int
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
