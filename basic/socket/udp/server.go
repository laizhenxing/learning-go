package main

import (
	"fmt"
	"net"
	"os"
)

// 通常用于视频直播
func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0,0,0,0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("upd listen server error: ", err)
		os.Exit(1)
	}
	defer listen.Close()

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])	// 接收数据
		if err != nil {
			fmt.Println("read upd failed, err: ", err)
			continue
		}
		fmt.Printf("data: %v, addr: %v, count: %d\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP([]byte("upd 发送数据"), addr)
		if err != nil {
			fmt.Println("udp 发送数据失败，err: ", err)
			continue
		}
	}
}
