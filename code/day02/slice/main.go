package main

import (
	"fmt"
	"sort"
)

func main() {

	// 切片定义
	var s1 []int // 定义一个存放 int 类型元素的切片
	var s2 []string
	println(s1, s2)        // [0/0]0x0 [0/0]0x0
	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) //true

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"shahe", "zhangjiang", "pingshancun"}
	fmt.Println(s1, s2)    // [1 2 3] [shahe zhangjiang pingshancun]
	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil) //false

	// 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4]   // 基于一个数组切割，左包含右不包含（左闭右开）
	fmt.Println(s3) // [1 3 5 7]
	s4 := a1[1:6]
	fmt.Println(s4) // [3 5 7 9 11]
	s5 := a1[:4]    // [1 3 5 7]
	s6 := a1[3:]    // [7 9 11 13]
	s7 := a1[:]     // [1 3 5 7 9 11 13]
	fmt.Println(s5, s6, s7)

	// 切片的容量：指底层数组的容量
	// 切片的容量：底层数组从切片的第一个元素到最后的元素数量
	// 切片的长度：切割的元素个数
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))

	// 切片再切割
	s8 := s6[3:]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))
	// 切片是引用类型，都指向了底层的一个数组
	fmt.Println("s6:", s6)
	a1[6] = 100
	fmt.Println("s6:", s6)
	fmt.Println("s8:", s8)

	// make 函数创造切片
	s9 := make([]int, 5, 10)
	fmt.Printf("s9:%v len(s9):%d cap(s9):%d\n", s9, len(s9), cap(s9))

	s10 := make([]int, 0, 10)
	fmt.Printf("s10:%v len(s10):%d cap(s10):%d\n", s10, len(s10), cap(s10))

	// 切片的赋值
	s11 := []int{1, 3, 5}
	s12 := s11 // s11 和 s12 都指向了同一个底层数组
	s11[0] = 100
	fmt.Println(s11, s12)

	// 切片的遍历
	// 1. 索引遍历
	for i := 0; i < len(s12); i++ {
		fmt.Println(s12[i])
	}

	// 2. for range 遍历
	for _, v := range s12 {
		fmt.Println(v)
	}

	// append 为切片追加元素
	s13 := []string{"beijing", "shanggai", "shenzhen"}
	fmt.Printf("s13:%v len(s13):%d cap(s13):%d\n", s13, len(s13), cap(s13))
	// s13[3] = "广州" // 编译报错，索引越界
	// fmt.Println(s13)

	// 调用 append 函数必须用原来的切片变量接收返回值，append 追加元素，原来的数组放不下爱的时候 go 底层会把底层数组换一个
	// 必须用变量接收 append 的返回值
	s13 = append(s13, "guangzhou")
	fmt.Printf("s13:%v len(s13):%d cap(s13):%d\n", s13, len(s13), cap(s13))
	s13 = append(s13, "hanghzou", "chengdu")
	fmt.Printf("s13:%v len(s13):%d cap(s13):%d\n", s13, len(s13), cap(s13))
	ss := []string{"huhan", "xian", "suzhou"}
	s13 = append(s13, ss...) // ... 表示把切片拆开
	fmt.Printf("s13:%v len(s13):%d cap(s13):%d\n", s13, len(s13), cap(s13))

	// copy
	aa1 := []int{1, 3, 5}
	aa2 := aa1
	var aa3 = make([]int, 3, 3)
	copy(aa3, aa1)
	fmt.Println(aa1, aa2, aa3) // [1 3 5] [1 3 5] [1 3 5]

	aa1[0] = 100
	fmt.Println(aa1, aa2, aa3) // [100 3 5] [100 3 5] [1 3 5]

	// 将aa1 中索引为1的3这个元素删掉
	aa1 = append(aa1[:1], aa1[2])
	fmt.Println(aa1, cap(aa1)) // [100 5] 3

	x1 := [...]int{1, 3, 5}
	s14 := x1[:] // [1 3 5]
	fmt.Println(s14, len(s14), cap(s14))

	// 切片不保存具体的值  切片对应一个底层数组  底层数组都占用一块连续的内存
	fmt.Printf("%p\n", &s14[0])
	s14 = append(s14[:1], s14[2:]...) // 修改了底层数组
	fmt.Printf("%p\n", &s14[0])
	fmt.Println(s14, len(s14), cap(s14))
	fmt.Println(x1) // [1 5 5]
	s14[0] = 100
	fmt.Println(x1, s14)

	// 切片练习题
	var aa4 = make([]int, 5, 10)
	fmt.Println(aa4)
	for i := 0; i < 10; i++ {
		aa4 = append(aa4, i)
	}
	fmt.Println(aa4)      // [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(cap(aa4)) // 20 append 底层数组容量不够的时候会自动换一个容量更大的底层数组

	// 对切片进行排序
	var aa5 = [...]int{1, 3, 5, 4, 2, 11}
	sort.Ints(aa5[:])
	fmt.Println(aa5) // [1 2 3 4 5 11]

	// 使用 append 删除切片中的某个元素
	aa6 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	ss15 := aa6[:]

	// 删掉索引为1的那个3
	ss15 = append(ss15[:1], ss15[2:]...)
	fmt.Println(ss15, aa6) // [1 5 7 9 11 13 15 17] [1 5 7 9 11 13 15 17 17]
}
