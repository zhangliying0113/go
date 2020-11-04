package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet uint64
}

func (c cat) move() {
	fmt.Println("走猫步。。。")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

type chicken struct {
	feet int64
}

func (c chicken) move() {
	fmt.Println("鸡动。。。")
}

func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s。。。\n", food)
}

func main() {
	var a1 animal          // 定义一个接口类型变量
	fmt.Printf("%T\n", a1) // <nil>

	bc := cat{
		name: "xiaohua",
		feet: 4,
	}

	a1 = bc
	a1.eat("小鱼")           // 猫吃小鱼...
	fmt.Printf("%T\n", a1) // main.cat
	fmt.Println(a1)        // {xiaohua 4}

	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	fmt.Printf("%T\n", a1) // main.chicken
}
