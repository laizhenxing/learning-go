package server

import (
	"bufio"
	"fmt"
	"net"
)

// TCP 服务端的处理流程
// 1. 监听端口
// 2. 接受客户端请求建立连接
// 3. 创建goroutine处理连接

func Process(conn net.Conn)  {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])	// 读取数据
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client 端发来的数据：", recvStr)
		rspStr := "你发来的消息是：" + recvStr
		conn.Write([]byte(rspStr))	// 发送数据,返回给客户端
	}
}