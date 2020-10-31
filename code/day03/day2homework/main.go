package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 1. 判断字符串中汉字的数量
	s1 := "Hello沙河안녕하세요こんにちは"
	var count int
	for _, v := range s1 {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println(count)

	// 2. how do you do 单词出现的次数
	s2 := "how do you do"
	s3 := strings.Split(s2, " ")
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	// 3. 会问判断 字符串从左往右读和从右往左读是一样的，那么就是回文。
	ss := "a山西运煤车煤运西山a"
	r := make([]rune, 0, len(ss))
	for _, v := range ss {
		r = append(r, v)
	}
	fmt.Println("r:", r)

	for i := 0; i < len(r)/2; i++ {
		if r[1] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")

}
