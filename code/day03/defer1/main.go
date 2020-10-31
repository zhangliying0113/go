package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

// 执行过程
// 1. a = 1  b = 2
// 2. defer calc("1", 1, calc("10", a, b))
// 3. calc("10", a, b) = "10" 1 2 3
// 4. defer calc("1", 1, 3)
// 5. a = 0
// 6. defer calc("2", a, calc("20", 0, 2))
// 7. calc("20", 0, 2) = "20" 0 2 2
// 8. defer calc("2", 0, 2)
// 9. b = 1
// 10. defer calc("2", 0, 2) = "2" 0 2 2
// 11. defer calc("1", 1, 3) = "1" 1 3 4

// 执行结果：
// 10 1 2 3
// 20 0 2 2
// 2 0 2 2
// 1 1 3 4
