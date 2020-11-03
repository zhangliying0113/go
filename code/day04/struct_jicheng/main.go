package main

import "fmt"

// 结构体模拟实现其他语言中的“继承”

type animal struct {
	name string
}

// 给 animal 实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

// dog 结构体
type dog struct {
	feet   uint8
	animal // animal 拥有的方法，dog 此时也有了
}

// 给 dog 实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s会叫！\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{
			name: "xiaohua",
		},
		feet: 4,
	}
	fmt.Println(d1) // {4 {xiaohua}}
	d1.wang()       // xiaohua会叫！
	d1.move()       // xiaohua会动！
}
