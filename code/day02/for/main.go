package main

import "fmt"

func main() {
	// 跳出 for 循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// 跳过此次 for 循环，继续执行下一次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
