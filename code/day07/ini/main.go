package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini配置文件解析器

// MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig ...
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}

// Config ...
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0. 参数校验
	t := reflect.TypeOf(data)
	// fmt.Println(t, t.Kind()) // *main.Config ptr 类型是结构体  种类是指针
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer") // 新建一个错误
		return
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a pointer") // 新建一个错误
		return
	}

	// 1. 读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	// 将字节类型的文件内容转化成字符串
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice) // string 类型切片
	// []string{"; mysql config", "[mysql]", "address=10.20.30.40", "port=3306", "username=root", "password=rootroot", "", "         ", "", "", " # redis config", ";  [", "; [  ]", "[redis]", "; heihei", "; = hahaha", "xxx=", "host =11.22.33.44", "port=6379", "password=root123", "database=0", "test=false"}

	// 2. 遍历切片，读出每一行数据
	var structName string
	for idx, line := range lineSlice {
		// 去掉字符串首尾的空格
		line = strings.TrimSpace(line)
		// 如果是空行就跳过
		if len(line) == 0 {
			continue
		}
		// 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 如果是[开头就是节(section)，对应 Config 结构体元素 MysqlConfig 或 RedisConfig
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 把这一行首尾的[]去掉，取到中间的内容把首尾的空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				// fmt.Println(field) // {MysqlConfig  main.MysqlConfig ini:"mysql" 0 [0] true}    {RedisConfig  main.RedisConfig ini:"redis" 56 [1] true}
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}

		} else {
			// 如果不是[开头就是=分割的键值对
			// 以等号分割这一行，等号左边是 key,等号右边是 value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// 根据 strucrName 去 data 里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型信息
			// fmt.Println(sType, sType.Kind())  // main.RedisConfig struct
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}

			// 遍历嵌套结构体的每一个字段，判断tag是不是等于key
			var fieldName string
			var fieldType reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) // // tag信息是存储在类型信息中的
				fieldType = field
				// fmt.Println(fieldType) // {Address  string ini:"address" 0 [0] false}
				if field.Tag.Get("ini") == key {
					fieldName = field.Name
					break
				}
			}
			if len(fieldName) == 0 {
				// 在结构体中找不到对应的字符
				continue
			}
			// 如果key = tag，给这个字段赋值
			// 根据fieldName 去取出这个字段
			fileObj := sValue.FieldByName(fieldName)
			// 对其赋值
			fmt.Println(fieldName, fieldType.Type.Kind())
			switch fieldType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}
		}
	}
	return
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", cfg)
}
