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

func (c *cat) move() {
	fmt.Println("走猫步。。。")
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	var a1 animal          // 定义一个接口类型变量
	fmt.Printf("%T\n", a1) // <nil>

	bc := cat{
		name: "xiaohua",
		feet: 4,
	}

	bd := &cat{
		name: "xiaobai",
		feet: 4,
	}

	a1 = &bc     // 实现animal这个接口的是cat的指针类型
	a1.eat("小鱼") // 猫吃小鱼...
	a1 = bd
	a1.move()
	fmt.Printf("%T\n", a1) // *main.cat
	fmt.Println(a1)        // &{xiaobai 4}
}
