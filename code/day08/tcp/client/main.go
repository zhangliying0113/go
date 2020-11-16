package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. 与 server 端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:20000 failed,err:", err)
		return
	}

	// 2. 发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("请说话：")
		msg, _ := reader.ReadString('\n')
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
