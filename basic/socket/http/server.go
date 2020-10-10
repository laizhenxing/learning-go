package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go", myHandler)

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RemoteAddr, "连接成功")
	fmt.Println(r.Method, "连接方式")
	fmt.Println(r.URL.Path, "请求地址")
	fmt.Println(r.Header, "请求头")
	fmt.Println(r.Body, "请求体")
	// 回复
	w.Write([]byte("服务器回复信息"))
}