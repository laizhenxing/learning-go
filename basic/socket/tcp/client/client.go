package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

func DefaultClient()  {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("与服务器建立连接失败。 err: ", err)
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')	// 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {	// 如果输入q就推出
			return
		}
		_, err =  conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("接收服务端信息失败. err: ", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func DealWithWaitGroup()  {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("connect to tcp server error: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)

	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup)  {
	defer wg.Done()

	for i := 10; i > 0; i-- {
		_, err := conn.Write([]byte("hello " + strconv.Itoa(i) + "\r\n"))
		if err != nil {
			fmt.Println("Error to send message because of ", err.Error())
			break
		}
	}
}

func handleRead(conn net.Conn, wg *sync.WaitGroup)  {
	defer wg.Done()

	reader := bufio.NewReader(conn)
	for i := 0; i < 10; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error to read message because of ", err.Error())
			return
		}
		fmt.Println(line)
	}
}