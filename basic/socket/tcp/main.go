package main

import (
	"fmt"
	"net"
	"os"

	"socket/tcp/server"
)

// server main
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed. err: ", err)
		os.Exit(1)
	}

	for {
		conn, err := listen.Accept()	// 建立连接
		if err != nil {
			fmt.Println("accept failed. err: ", err)
			continue
		}

		// logs an incoming message
		fmt.Printf("Recieved message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		go server.Process(conn)	// 启动一个 goroutine 处理连接
	}
}
