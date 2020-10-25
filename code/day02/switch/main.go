package main

import "fmt"

func main() {
	// 情况1
	var n = 3
	switch n {
	case 1:
		fmt.Println("is 1")
	case 2:
		fmt.Println("is 2")
	case 3:
		fmt.Println("is 3")
	case 4:
		fmt.Println("is 4")
	case 5:
		fmt.Println("is 5")
	}

	// 情况2
	switch n := 3; n {
	case 1:
		fmt.Println("is 1")
	case 2:
		fmt.Println("is 2")
	case 3:
		fmt.Println("is 3")
	case 4:
		fmt.Println("is 4")
	case 5:
		fmt.Println("is 5")
	default:
		fmt.Println("unabled")
	}

	// 情况3
	switch n := 10; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("is jishu")
	case 2, 4, 6, 8:
		fmt.Println("is oushu")
	default:
		fmt.Println(n)
	}
}
