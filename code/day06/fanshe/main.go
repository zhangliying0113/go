package main

import (
	"fmt"
	"reflect"
)

type cat struct {
}

type animal struct {
	Name string `json:"name" xiaobai:"bai"`
	Age  int    `json:"age" xiaobai:"hei"`
}

// 获取变量类型和种类
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v type kind:%v\n", v.Name(), v.Kind())
}

// 获取变量值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind() // 值得类型种类
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取整型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取整型的原始值，然后通过float64强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置变量的值
func refelectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本，reflect 包会引发 panic
	}
}

func refelectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) // 修改的是副本，reflect 包会引发 panic
	}
}

func main() {
	// float32
	// type:float32
	// type name:float32 type kind:float32
	var a float32 = 3.14
	reflectType(a)
	reflectValue(a) // type is float32, value is 3.140000

	// int64
	// type:int64
	// type name:int64 type kind:int64
	var b int64 = 100
	reflectType(b)

	// main.cat
	// type:main.cat
	// type name:cat type kind:struct
	var c = cat{}
	reflectType(c)

	// refelectSetValue1(b)   // panic: reflect: reflect.Value.SetInt using unaddressable value

	refelectSetValue1(&b) // 传入地址没有 panic，但是修改不成功，需要用 Elem才可以
	fmt.Println(b)        // 100
	refelectSetValue2(&b)
	fmt.Println(b) // 200

	a1 := animal{
		Name: "xiaobai",
		Age:  1,
	}

	t := reflect.TypeOf(a1)
	fmt.Println(t.Name(), t.Kind()) // animal struct

	// for 循环遍历结构体的所有字段信息
	fmt.Println(t.NumField()) // 2
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json")) // name:Name index:[0] type:string json tag:name
	}

	// 根据字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("xiaobai")) // name:Age index:[1] type:int json tag:age
	}
}
