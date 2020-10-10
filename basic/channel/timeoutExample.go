package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "result"
	}()

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timeout")
	case res := <-ch:
		fmt.Println(res)
	}
}
