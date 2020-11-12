package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	name string
	city string
	age  int64
}

type inter struct {
	a int8
	b int8
	c int8
}

func main() {
	var p1 person
	fmt.Printf("p1=%v\n", &p1)
	fmt.Printf("p1=%#v\n", p1)
	p1.name = "小白"
	p1.age = 2
	fmt.Printf("p1=%#v\n", p1)

	var p2 = new(person)
	p2.name = "小白"
	p2.age = 18
	fmt.Printf("%T\n", p2)
	fmt.Printf("p2=%#v\n", p2)

	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)

	p4 := person{
		name: "小白",
		city: "天津",
		age:  2,
	}
	fmt.Printf("p4=%#v\n", p4)

	p5 := &person{
		name: "小白",
		city: "天津",
		age:  2,
	}
	fmt.Printf("p5=%#v\n", p5)

	// p6 := &person{
	// 	"小白",
	// 	"天津",
	// 	2,
	// }

	i := inter{
		1, 2, 3,
	}
	fmt.Printf("i.a=%#v\n", &i.a)
	fmt.Printf("i.b=%#v\n", &i.b)
	fmt.Printf("i.c=%#v\n", &i.c)

	var ss struct{}
	fmt.Printf("ss%#v\n", &ss)
	fmt.Println(unsafe.Sizeof(ss))
}
