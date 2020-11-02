package main

import "fmt"

// 递归：函数自己调用自己
// 递归适合处理：问题相同且问题的规模越来越小的场景
// 递归一定要有一个明确的退出条件

// 实例1：计算 n 的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

// 实例2：上台阶问题
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return taijie(n-1) + taijie(n-2)
}

func main() {
	ret := f(7)
	fmt.Println(ret) // 5040
	ret1 := taijie(4)
	fmt.Println(ret1) // 5
}
