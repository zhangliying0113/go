package main

import "fmt"

func main() {
	var name string // 变量声明
	name = "理想"
	fmt.Println(name)
	// 数组
	var ages [30]int // 数组声明，声明了一个变量 ages，他是 [30]int 类型
	ages = [30]int{1, 2, 3, 4, 5, 6}
	fmt.Println(ages) // [1 2 3 4 5 6 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	var ages2 = [...]int{1, 2, 3, 4, 5, 7, 8, 90}
	fmt.Println(ages2)
	var ages3 = [...]int{1: 20, 99: 200}
	fmt.Println(ages3) // [0 20 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 200]
	// var ages4 = [...]int{"tom": 20, "jerry": 200} // 写法错误，数组索引值必须为 int 型常量

	// 二维数组
	// 多维数组只有最外层可以使用...
	var a1 = [...][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a1) // [[1 2] [3 4] [5 6]]

	x := [3]int{1, 2, 3}
	y := x     // 把 x 的值拷贝了一份给 y
	y[1] = 200 // 修改的是副本 y，并不影响 x

	fmt.Println(x) // [1 2 3]
	fmt.Println(y) // [1 200 3]
	f1(x)
	fmt.Println(x) // [1 2 3]

	// 切片
	var s1 []int           // 没有分配内存，== nil
	fmt.Println(s1)        // []
	fmt.Println(s1 == nil) // true
	s1 = []int{1, 2, 3}
	fmt.Println(s1) // [1 2 3]

	// make 初始化分配内存
	s2 := make([]bool, 2, 4)
	fmt.Println(s2) // [false false]
	s3 := make([]int, 0, 4)
	fmt.Println(s3 == nil) // false

	s4 := []int{1, 2, 3}
	s5 := s4
	var s6 = make([]int, 3, 3)
	copy(s6, s4)
	fmt.Println(s5) // [1 2 3]
	s5[1] = 200
	fmt.Println(s5) // [1 200 3]
	fmt.Println(s4) // [1 200 3] // 切片是引用类型， s4 和 s5 指向同一内存空间
	fmt.Println(s6) // [1 2 3] // copy 得到的 s6 相当于副本，不会发生改变

	var s7 []int // nil, 未分配内存
	s7 = make([]int, 1)
	s7[0] = 100
	fmt.Println(s7)    // [100]
	s7 = append(s7, 1) // 自动扩容
	fmt.Println(s7)    // [100 1]

	// 指针
	// go 语言中指针只能读不能修改，不能修改指针变量指向的地址
	addr := "沙河"
	addrP := &addr
	fmt.Println(addrP)        // 内存地址 0xc00008c000
	fmt.Printf("%T\n", addrP) // *string
	addrV := *addrP           // 根据内存地址找值
	fmt.Println(addrV)        // 沙河

	// map
	var m1 map[string]int
	fmt.Println(m1 == nil) // true 未开辟内存
	m1 = make(map[string]int, 10)
	m1["理想"] = 100
	fmt.Println(m1)       // map[理想:100]
	fmt.Println(m1["ji"]) // 0 如果 key 不存在返回的是 value 对应类型的零值
	// 如果返回值是布尔型，我们通常用 ok 去接收他
	score, ok := m1["ji"]
	if !ok {
		fmt.Println("没有姬无命这个人")
	} else {
		fmt.Println("姬无命的分数是：", score)
	}

	delete(m1, "lixiang") // 删除的 key 不存在什么操作都不做
	delete(m1, "理想")
	fmt.Println(m1)        // map[]
	fmt.Println(m1 == nil) // false 开辟了内存

}
func f1(a [3]int) {
	// go 语言中的函数传递的都是值
	a[1] = 100 // 此处修改的是副本的值
}
