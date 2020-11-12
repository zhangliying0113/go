package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}

	// 设置往文件写入
	log.SetOutput(fileObj) // 默认 os.Stdout 从终端输出

	for {
		log.Println("这是一条测试的日志") // 2020/11/09 13:12:36 这是一条测试的日志   打印的时候会直接把时间打印上，没有开关和区分日志级别
		time.Sleep(time.Second)
	}

}
