package main

import "fmt"

func main() {
	// 数组：指定存放的元素类型和容量，数组的长度是数组类型的一部分
	// 数组声明
	// 如果不初始化：默认元素都是零值（布尔值：false, 整型和浮点型都是0, 字符串：""）
	var a1 [3]int  // 默认值[0 0 0]
	var a2 [4]bool // 默认值[false false false false]
	fmt.Printf("a1:%T--%v\n", a1, a1)
	fmt.Printf("a2:%T--%v\n", a2, a2)

	// 数组初始化
	// 1. 初始化方式1
	a1 = [3]int{1, 2, 3}
	a2 = [4]bool{true, false, true, true}

	// 2. 初始化方式2：根据初始值自动推断数组的长度是多少
	a10 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a10)

	// 3. 初始化方式3：根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 数组遍历
	citys := [...]string{"北京", "上海", "深圳"}

	// 1. 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	// 2. for range 遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 二维数组

	// 声明
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	a12 := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)
	fmt.Println(a12)

	// 多维数组遍历
	for _, v := range a11 {
		fmt.Println(v)
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)

	// 数组练习题
	// 1. 求数组[1, 3, 5, 7, 8]所有元素的和
	c1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range c1 {
		sum += v
	}
	fmt.Println(sum)

	// 2. 找出数组中和为指定值(8)的两个元素的下标
	for i := 0; i < len(c1); i++ {
		for j := i + 1; j < len(c1); j++ {
			if c1[i]+c1[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
