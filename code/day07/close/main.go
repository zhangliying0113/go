package main

import "fmt"

func main() {
	ch1 := make(chan bool, 2)
	ch1 <- true
	ch1 <- true
	close(ch1)
	// for x := range ch1 {
	// 	fmt.Println(x) // 通道管理之后里面数据仍然存在，这里拿值之后后面再从通道取值无值，返回零值 false
	// }

	// <-ch1
	// <-ch1
	x, ok := <-ch1
	fmt.Println(x, ok) // true true
	x, ok = <-ch1
	fmt.Println(x, ok) // true true
	x, ok = <-ch1
	fmt.Println(x, ok) // false false
}
