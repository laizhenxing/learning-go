package main

import "socket/tcp/client"

// TCP客户端进行TCP通信的流程
// 1. 建立与服务端的连接
// 2. 进行数据收发
// 3. 关闭连接
func main() {
	//client.DealWithWaitGroup()
	client.DefaultClient()
}
