// Fetchall  fetches URLs in parallel and reports their times and sizes
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)	// start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)	// receive from channel ch
	}
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	rsp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, rsp.Body)
	// 关闭资源
	rsp.Body.Close()
	if  err != nil {
		ch <- fmt.Sprintf("while reading %s : %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
