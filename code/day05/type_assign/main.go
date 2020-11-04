package main

import "fmt"

type ff func()

// 类型断言1
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串：", str)
	}
}

// 类型断言2
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("传进来的是一个字符串：", t)
	case int:
		fmt.Println("传进来的是一个int：", t)
	case int64:
		fmt.Println("传进来的是一个int64：", t)
	case []int:
		fmt.Println("传进来的是一个slice：", t)
	case map[string]int:
		fmt.Println("传进来的是一个map：", t)
	case func():
		fmt.Println("传进来的是一个func：", t) // 传入了函数类型但是未被调用，会提示绿线但是可以编译通过
	}
}

func f() {

}

func main() {
	assign(100)                     // int 猜错了
	assign2(true)                   // bool
	assign2(int64(200))             //int64 传进来的是一个int64： 200
	assign2([]int{1, 2, 3})         // []int 传进来的是一个slice： [1 2 3]
	assign2(map[string]int{"a": 1}) // map[string]int 传进来的是一个map： map[a:1]
	assign2(f)                      // func() 传进来的是一个func： 0xad79c0,返回的是地址
}
