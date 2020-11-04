package main

import "fmt"

// 接口实例 定义一个能叫的类型
type speaker interface {
	speak() // 只要实现了 speak 方法的变量都是 speaker 类型，方法签名
}

type cat struct{}
type dog struct{}
type person struct{}

func (c cat) speak() {
	fmt.Println("miaomiaomiao")
}

func (d dog) speak() {
	fmt.Println("wangwangwang")
}

func (p person) speak() {
	fmt.Println("lalala")
}

func da(s speaker) {
	s.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1) // miaomiaomiao
	da(d1) // wangwangwang
	da(p1) // lalala

	var ss speaker // 定义一个接口类型 speaker 的变量 ss
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss) // {}
}
