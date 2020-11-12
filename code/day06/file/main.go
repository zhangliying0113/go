package main

import (
	"fmt"
	"io"
	"os"
)

// 文件操作
// 文件中间写入内容

func f1() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./xxx.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
	}
	// defer 不能写在 if 语句上面，若 err 不等于 nil，则 fileObj 为指针的零值 nil，nil.Close() 可能引发 panic，而文件函数 Close 不成功会返回 err，不会 panic，但是其他情况可能会。
	defer fileObj.Close()
}

func f2() {
	// 打开要操作的文件
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed, err:", err)
	}
	defer fileObj.Close()

	// 因为没有办法直接在文件中间插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./sb.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open tmp file failed, err:", err)
	}
	defer tmpFile.Close()

	// 读取源文件写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Println("read sb file failed,err:", err)
	}

	// 写入临时文件
	tmpFile.Write(ret[:n])

	// 在临时文件写入要插入的内容
	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)

	// 再把源文件后续的内容读出写入到临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Println("read sb file failed,err:", err)
			return
		}

		tmpFile.Write(x[:n])
	}

	// 源文件也后续的写入了临时文件中
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./sb.tmp", "./sb.txt")
}

func main() {
	f2()
}
