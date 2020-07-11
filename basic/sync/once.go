package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	once := sync.Once{}

	for i := 0; i < 1000; i++ {
		go func(idx int) {
			once.Do(func() {
				time.Sleep(time.Second)
				fmt.Println("hello world~ index: ", idx)
			})
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("main func done~")
}
