package main

import "fmt"

// 方法

// 标识符：变量名 函数名 类型名 方法名
// go 语言中如果标识符首字母是大写的，就表示对外部包可见(暴露的，共有的)

type dog struct {
	name string
}

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~\n", d.name)
}

// 使用值接收者：传拷贝进去
func (p person) guonian() {
	p.age++
}

// 使用指针接收者：传内存地址进去
func (p *person) zhenguonian() {
	p.age++
}

func (p *person) dream() {
	fmt.Println("不上班也能挣钱")
}

// 给自定义类型加方法
// 不能给别的包里面的类型添加方法，只能给自己包里的类型添加方法

type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个 int")
}
func main() {
	d1 := newDog("erha")
	d1.wang()

	p1 := newPerson("xianger", 18)
	// p1.wang() 报错，wang 的接收者为 dog 类型
	p1.guonian()
	fmt.Println(p1.age) // 18 值传递

	p1.zhenguonian()
	fmt.Println(p1.age) // 19 地址传递

	p1.dream()

	// 调用自定义类型的方法
	m := myInt(100)
	fmt.Printf("%T  %v\n", m, m) // main.myInt  100
	m.hello()
}
