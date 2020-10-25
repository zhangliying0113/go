package main

import (
	"fmt"
	"strings"
)

// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。

func main() {
	str := "how do you do"
	fmt.Println(count(str))
}

func count(str string) (ret map[string]int) {
	count := make(map[string]int)
	str2 := strings.Split(str, " ")
	for _, value := range str2 {
		count[value]++
	}
	return count
}

//make 和 new
// make和new都是用来申请内存的
// new很少用，一般用来给基本数据类型申请内存，`string`、`int`,返回的是对应类型的指针(\*string、\*int)。
// make是用来给`slice`、`map`、`chan`申请内存的，make函数返回的的是对应的这三个类型本身
