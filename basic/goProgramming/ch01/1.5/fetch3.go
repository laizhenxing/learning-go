// Fetch 3 add prefix(`http://`) for url
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// using strings.HasPrefix to check the url's prefix
		if b := strings.HasPrefix(url, "http://"); !b {
			url = "http://" + url
		}

		rsp, err := http.Get(url)
		fmt.Println("status: ", rsp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch3: %v\n", err)
			os.Exit(1)
		}
		dst := new(bytes.Buffer)
		res, err := io.Copy(dst, rsp.Body)
		// 关闭资源
		rsp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch3: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		//fmt.Println(dst)
		fmt.Println("response bytes: ", res)
	}
}
