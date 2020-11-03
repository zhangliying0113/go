package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与 json

// 序列化：把 go 语言中的结构体变量 --》 json 格式的字符串
// 反序列化： json 格式的字符串 --》 go 语言中能够识别的结构体变量

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "xiaohua",
		Age:  20,
	}

	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b)) // {"Name":"xiaohua","Age":20} 加了`json:"name"，序列化后的变量名会变 {"name":"xiaohua","age":20}

	// 反序列化
	str := `{"Name":"xiaohua","Age":20}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) // 传地址是为了能在 json.Unmarshal 内部修改 p2 的值
	fmt.Printf("%#v\n", p2)          // main.person{Name:"xiaohua", Age:20}
}
