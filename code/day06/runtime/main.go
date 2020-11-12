package main

import (
	"fmt"
	"path"
	"runtime"
)

func f1() {
	pc, file, line, ok := runtime.Caller(0) // 第几层调用
	if !ok {
		fmt.Printf("runtime.caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)        // main.f1 拿到运行的函数名，目前拿到的不是
	fmt.Println(file)            // D:/gocode/src/github.com/zhangliying0113/go/code/day06/runtime/main.go  一层一层的调用 参数传的是调用的第几层
	fmt.Println(line)            // 9 拿到当前层调用的行号
	fmt.Println(path.Base(file)) // main.go 拿到文件名
}

func main() {
	f1()
}
