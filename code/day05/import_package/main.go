package main

import (
	"fmt"

	calc "github.com/zhangliying0113/go/code/day05/calc"
)

var x = 100

const pi = 3.14

func init() {
	fmt.Println("自动执行。。。")
	fmt.Println(x, pi)
}

func main() {
	ret := calc.Add(100, 200)
	fmt.Println(ret)
}
