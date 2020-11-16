package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/zhangliying0113/go/code/day08/jiejue_zhanbao/protocol"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		recvStr, err := protocol.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode failed,err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", recvStr)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
