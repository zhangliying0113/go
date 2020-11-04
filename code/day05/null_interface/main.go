package main

import "fmt"

// 空接口类型 interface{}

func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "xiaobai"
	m1["age"] = 1
	m1["merried"] = true
	m1["hobby"] = []string{
		"唱歌",
		"tiaowu",
		"daqiu",
	}

	show(false) // type:bool value:false
	show(nil)   //type:<nil> value:<nil>
	show(m1)    // type:map[string]interface {} value:map[age:1 hobby:[唱歌 tiaowu daqiu] merried:true name:xiaobai]
}
