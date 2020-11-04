package main

import "fmt"

// 接口示例2
// 不管是什么牌子的车,都能跑

// 定义一个 car 的接口类型
type car interface {
	run()
}

type falali struct {
	brand string
}

type baoshijie struct {
	brand string
}

func (f falali) run() {
	fmt.Printf("%s速度70迈\n", f.brand)
}

func (b baoshijie) run() {
	fmt.Printf("%s速度100迈\n", b.brand)
}

func drive(c car) {
	c.run()
}

func main() {
	var f1 = falali{
		brand: "falali",
	}
	var b1 = baoshijie{
		brand: "baoshijie",
	}

	drive(f1) // falali速度70迈
	drive(b1) // baoshijie速度100迈
}
