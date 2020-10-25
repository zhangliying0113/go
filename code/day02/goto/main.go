package main

import "fmt"

func main() {
	// 跳出多层 for 循环
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break
			}
			fmt.Printf("%v--%c\n", i, j)
		}
		if flag {
			fmt.Println("over")
			break
		}
	}

	// goto+label 实现跳出多层 for 循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto ERROR
			}
			fmt.Printf("%v--%c\n", i, j)
		}
	}
ERROR:
	fmt.Println("over")
}
