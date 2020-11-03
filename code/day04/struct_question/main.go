package main

import "fmt"

// 1. myInt(100)是个啥?
type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

type person struct {
	name string
	age  int
}

func main() {
	// 声明一个 int32 类型的变量 x,它的值是10
	// 方法1:
	// var x int32
	// x = 10
	// 方法2:
	// var x int32 = 10
	// 方法3:
	// var x = int32(10)
	// // 方法4:
	// x := int32(10)
	// fmt.Println(x)

	// 声明一个 myInt 类型的变量 m,它的值是100
	// var m myInt
	// m = 100

	// var m myInt = 100

	// var m = myInt(100)

	m := myInt(100)
	m.hello()

	// 问题2：结构体初始化
	// 方法1:
	var p person // 声明一个 person 类型的变量 p
	p.name = "xianger"
	p.age = 800
	fmt.Println(p)

	// 方法2:
	s1 := []int{1,
		2,
		3,
		4,
	}
	m1 := map[string]int{
		"stu1": 100,
		"stu2": 200,
		"stu3": 300,
	}
	fmt.Println(s1, m1)
	// 键值对初始化
	var p2 = person{
		name: "xiaosi",
		age:  15,
	}
	fmt.Println(p2)
	// 值列表初始化
	var p3 = person{
		"wangwu",
		20,
	}
	fmt.Println(p3)

	// 匿名结构体字段练习
	st1 := student{
		"meiwen",
		25,
	}
	fmt.Println(st1)
	fmt.Println(st1.string)
	fmt.Println(st1.int)
}

// 问题3：为什么要有构造函数
// 别人调用我，我能给他一个 person 类型的变量
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

// 4. 匿名字段 字段比较少也比较简单的场景，不常用

type student struct {
	string
	int
}
