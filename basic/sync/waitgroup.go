package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			time.Sleep(time.Second)
			fmt.Println("hello world~")
		}()
	}
	// 等待所有协程结束
	wg.Wait()
	fmt.Println("WaitGroup all process done~")
}
