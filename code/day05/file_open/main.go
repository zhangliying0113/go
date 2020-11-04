package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 打开文件 按字节读取文件
func readFromFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}

	// 记得关闭文件
	defer fileObj.Close()

	// 读文件
	// 指定读的长度
	//var tmp = make([]byte, 128)
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}

	}
}

// 利用 bufio 这个包读取文件
func readFromFileByBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed err:%v", err)
		return
	}

	// 记得关闭文件
	defer fileObj.Close()

	// 创建一个用来从文件读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v", err)
			return
		}
		fmt.Print(line) // 此处不需要换行

	}
}

// 利用 "io/ioutil" 这个包读取文件,直接将整个文件内容输出
func readFromFileByIoutil() {

	// 直接读取整个文件内容
	fileObj, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("open file failed err:%v", err)
		return
	}

	fmt.Println(string(fileObj))
}

func main() {
	// readFromFile1()
	// readFromFileByBufio()
	readFromFileByIoutil()
}
