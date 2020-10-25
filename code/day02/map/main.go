package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// 声明
	var m1 map[string]int
	fmt.Println(m1 == nil) // true, 没有初始化 未在内存中开辟空间

	m1 = make(map[string]int, 10) // 需要估算好该 map 容量，避免在程序运行期间再动态扩容
	m1["hebei"] = 130
	m1["tianjin"] = 120
	fmt.Println(m1)          // map[hebei:130 tianjin:120]
	fmt.Println(m1["hebei"]) // 130

	// 约定俗称用 ok 接收返回的布尔值
	fmt.Println(m1["shijiazhuang"]) // 0 如果不存在这个 key 值拿到对应值类型的零值

	value, ok := m1["hebei"] // 130 true

	if !ok {
		fmt.Println("no have the key")
	} else {
		fmt.Println(value, ok)
	}

	// map 的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 只遍历 k
	for k := range m1 {
		fmt.Println(k)
	}

	// 只遍历v
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 删除
	delete(m1, "hebei")
	fmt.Println(m1)
	delete(m1, "shijiazhuang") // 删除不存在的 key,不会执行任何操作

	// 练习
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成 stu 开头的字符串
		value := rand.Intn(100)          // 生成0~99的随机数
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)

	// 取出 map 中的所有 key 存入切片 keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	fmt.Println(keys)

	// 对切片排序
	sort.Strings(keys)
	// 按照排序后的 key 遍历 map
	for _, v := range keys {
		fmt.Println(v, scoreMap[v])
	}

	// map 和 slice 组合
	// 元素类型为 map 的切片， 没有对内部的 map 做初始化
	var s1 = make([]map[int]string, 10, 10)
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "shahe"
	fmt.Println(s1) // [map[10:shahe] map[] map[] map[] map[] map[] map[] map[] map[] map[]]

	// 值为切片类型的 map
	var mm = make(map[string][]int, 10)
	mm["beijing"] = []int{10, 20, 30}
	fmt.Println(mm)

}
