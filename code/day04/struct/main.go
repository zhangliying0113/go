package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明一个 person 类型的变量 p
	var p person
	// 通过字段赋值
	p.name = "zhangxiang"
	p.age = 25
	p.gender = "nv"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p) // {zhangxiang 25 nv [篮球 足球 双色球]}

	// 访问变量 p 的字段
	fmt.Printf("%T\n", p) // main.person
	fmt.Println(p.name)   // zhangxiang

	var p2 person
	p2.name = "xyz"
	p2.age = 18
	fmt.Printf("type:%T value:%v\n", p2, p2) // type:main.person value:{xyz 18  []}

	// 匿名结构体，多用于临时场景
	var s struct {
		x string
		y int
	}

	s.x = "shenzhou"
	s.y = 80
	fmt.Printf("type:%T value:%v\n", s, s) // type:struct { x string; y int } value:{shenzhou 80}
}
