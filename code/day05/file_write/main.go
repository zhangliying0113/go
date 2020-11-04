package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件内容

func writeDemo1() {
	fileObj, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}

	defer fileObj.Close()
	// Write
	fileObj.Write([]byte("xiaobaixiaobai!\n"))

	// WriteString
	fileObj.WriteString("小白是小白兔")
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("小白 你好呀\n") //写到缓存中
	wr.Flush()                 // 将缓存中的内容写入文件
	fileObj.Close()
}

func writeDemo3() {
	str := "hello 小白"
	err := ioutil.WriteFile("test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	// writeDemo1()
	// writeDemo2()
	writeDemo3()
}
