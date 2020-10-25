package main

import "fmt"

func main() {
	// & 取地址
	n := 18
	p := &n
	fmt.Println(p)        // 0xc0000140a0 变量地址
	fmt.Printf("%T\n", p) // *int

	// * 根据地址取值
	m := *p
	fmt.Println(m)        // 18
	fmt.Printf("%T\n", m) // int

	var a1 *int // <nil>  未开辟内存地址
	fmt.Println(a1)
	var a2 = new(int) // new函数申请一个内存地址
	fmt.Println(a2)   // 0xc0000140e8
	fmt.Println(*a2)  // 根据地址取值 0
	*a2 = 100
	fmt.Println(*a2) // 根据地址取值 100
}
