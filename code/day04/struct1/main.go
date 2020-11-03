package main

import "fmt"

// 结构体是值类型

type person struct {
	name, gender string
}

// go 语言中函数传参永远传的是拷贝
func f(x person) {
	x.gender = "nv"
}

func f2(x *person) {
	//(*x).gender = "nv" // 根据内存地址找到那个原变量，修改的就是原来变量的值,go 语言语法糖可以简写成如下：
	x.gender = "nv"

}

// 结构体占用一块连续的内存空间
type x struct {
	a int8 // 8bit =》1byte
	b int8
	c int8
}

func main() {
	var p person
	p.name = "张香"
	p.gender = "nan"

	f(p) // nan 值拷贝，未改变外层数值
	fmt.Println(p.gender)

	f2(&p) // nv 传递地址，函数改变地址上的变量
	fmt.Println(p.gender)

	// 结构体指针1
	var p2 = new(person)
	(*p2).name = "lixiang"
	p2.gender = "zhong"

	fmt.Printf("%T\n", p2)  // *main.person p2 是指针类型
	fmt.Printf("%p\n", p2)  // 0xc0000044a0 p2 保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2) // 0xc000006030 求 p2 的内存地址

	// 结构体指针2
	// 情况1
	var p3 = &person{
		name: "xyz",
	}
	fmt.Printf("%#v\n", p3) // &main.person{name:"xyz", gender:""}  p3 是内存地址

	// 情况2
	p4 := &person{
		"小王子",
		"nan",
	}
	fmt.Printf("%#v\n", p4) // &main.person{name:"小王子", gender:"nan"} 同理，p4 是内存地址

	// 练习：
	var a int
	a = 100
	b := &a
	fmt.Printf("type a:%T type b:%T\n", a, b) // type a:int type b:*int
	// 将 a 的十六进制内存地址打印出来
	fmt.Printf("%p\n", &a) // 0xc000012108 a 的内存地址是 b
	fmt.Printf("%p\n", b)  // 0xc000012108
	fmt.Printf("%v\n", b)  // 0xc000012108
	fmt.Printf("%p\n", &b) // 0xc000006038 b 的内存地址

	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}

	fmt.Printf("%p\n", &(m.a)) // 0xc0000a20c8
	fmt.Printf("%p\n", &(m.b)) // 0xc0000a20c9
	fmt.Printf("%p\n", &(m.c)) // 0xc0000a20ca
}
