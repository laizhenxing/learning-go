package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		rsp, err := http.Get(url)	// 创建http请求
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(rsp.Body)
		rsp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
