package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口还可以嵌套

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eating(string)
}

type cat struct {
	name string
	feet int8
}

// cat 的指针类型实现了 mover 接口
func (c *cat) move() {
	fmt.Println("走猫步。。。")
}

// cat 的指针类型实现了 eater 接口
func (c *cat) eating(food string) {
	fmt.Printf("猫是%s...\n", food)
}

func main() {
	var a1 animal
	a1 = &cat{
		name: "xiaobai",
		feet: 4,
	}

	fmt.Printf("%T\n", a1) // *main.cat

	mover.move(a1) // 走猫步。。。  ? 怎么通过 animal 访问 move
}
