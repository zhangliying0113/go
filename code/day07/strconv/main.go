package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {

	// 从字符串中解析出对应的数据
	str := "10000"
	// ret := int64(str) 报错：cannot convert str (type string) to type int64
	ret, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("ParseInt failed, err:", err)
	}
	fmt.Printf("%#v %T\n", ret, ret)      // 10000 int64
	fmt.Printf("%#v %T\n", ret, int(ret)) // 10000 int

	// Atoi 字符串转成 int
	ret1, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", ret1, ret1) // 10000 int

	// 从字符串中解析出 bool 值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue) // true bool

	// 从字符串中解析出浮点数
	floatStr := "1.34zhangwang白"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue) // 0 float64 解析不成功，返回对应数据类型的零值

	floatStr1 := "1.34"
	floatValue1, _ := strconv.ParseFloat(floatStr1, 64)
	fmt.Printf("%#v %T\n", floatValue1, floatValue1) // 1.34 float64

	// 把数字转化成字符串类型
	i := 97
	ret2 := string(i)
	fmt.Printf("%#v %T\n", ret2, ret2) // "a" string  转化的结果是 ASC|| 码，对应字符，不是数字字符串

	ret3 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v %T\n", ret3, ret3) // "97" string

	ret4 := strconv.Itoa(i)
	fmt.Printf("%#v %T\n", ret4, ret4) // "97" string
}
