package main

import "fmt"

func main() {
	// if 条件判断
	age := 19
	if age > 18 {
		fmt.Println("成年啦")
	} else if age > 10 {
		fmt.Println("超过10岁啦")
	} else {
		fmt.Println("好好学习")
	}

	// 作用域问题 newAge 此时只在 if 条件判断中生效
	if newAge := 10; newAge > 18 {
		fmt.Println("成年啦")
	} else {
		fmt.Println("好好学习")
	}

	// for 循环

	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 变种1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// 变种2
	var j = 3
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// // 无限循环
	// for {
	// 	fmt.Println(5)
	// }

	// for range 循环
	s := "hello 沙河"
	for _, v := range s {
		fmt.Printf("%c\n", v) // v 为字符对应索引值
	}

	for k, v := range s {
		fmt.Printf("%d %c\n", k, v)
	}
}
