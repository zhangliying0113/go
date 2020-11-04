package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时如果有空格

func useScan() {
	var s string
	fmt.Print("please input：")
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是%s\n", s)
	// lease input：12345 6789
	// 你输入的内容是12345 空格及之后内容无法获取
}

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please input：")
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是%s\n", s)
	// please input：1234 56798  00
	// 你输入的内容是1234 56798  00

}

func main() {
	// useScan()
	// useBufio()

	// 写日志
	fmt.Fprintln(os.Stdout, "这是一条日志记录")
	fileObj, _ := os.OpenFile("./rizhi.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	fmt.Fprintln(fileObj, "这是一条日志")
}
