package main

import (
	"fmt"
	"os"
)

// 1. 获取文件对象的类型
// 2. 获取文件对象的详细信息

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 1. 获取文件对象的类型
	fmt.Printf("%T\n", fileObj) // *os.File

	// 2. 获取文件对象的详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return
	}

	fmt.Printf("file size is %dB\n", fileInfo.Size()) // file size is 596B
	fmt.Printf("file name is %s\n", fileInfo.Name())  // file name is main.go  返回文件名，不包含路径
}
