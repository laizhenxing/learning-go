package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	rsp, _ := http.Get("http://127.0.0.1:8080/go")
	defer rsp.Body.Close()

	fmt.Println(rsp.Status)
	fmt.Println(rsp.Header)

	buf := make([]byte, 1024)
	for {
		// 接收服务器信息
		n, err := rsp.Body.Read(buf)
		if err != nil || err != io.EOF {
			fmt.Println(err)
			return
		}
		fmt.Println("读取完毕")
		res := string(buf[:n])
		fmt.Println(res)
		break
	}
}