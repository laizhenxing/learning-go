package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0,0,0,0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("udp 客户端监听失败，err: ", err)
		os.Exit(1)
	}
	defer socket.Close()

	sendData := []byte("Hello udp server")
	_, err = socket.Write(sendData)	// 发送数据
	if err != nil {
		fmt.Println("udp客户端发送数据失败，err: ", err)
		os.Exit(1)
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data)	// 接收数据
	if err != nil {
		fmt.Println("udp客户端接收数据失败， err: ", err)
		os.Exit(1)
	}
	fmt.Printf("recv: %v, addr: %v, count: %d\n", string(data[:n]), remoteAddr, n)
}