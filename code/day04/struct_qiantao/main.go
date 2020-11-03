package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	address
	workPlace workPlace
}

type company struct {
	name string
	address
}

func main() {
	p1 := person{
		name: "xiaowang",
		age:  30,
		address: address{
			province: "hebei",
			city:     "shijiahzhuang",
		},
	}
	fmt.Println(p1) // {xiaowang 30 {hebei shijiahzhuang} { }}  workPlace没复制默认为空结构体

	fmt.Println(p1.name, p1.address, p1.address.city) // xiaowang {hebei shijiahzhuang} shijiahzhuang
	fmt.Println(p1.city)                              // shijiahzhuang 先在自己结构体找这个字段,找不到就去匿名嵌套的结构体中查找该字段，不会找有名字的结构体
	fmt.Println(p1.address.city)                      //shijiahzhuang
	fmt.Println(p1.workPlace.city)                    // 没赋值，为空
}
