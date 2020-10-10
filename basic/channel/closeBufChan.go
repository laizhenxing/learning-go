package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 20)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	go func() {
		for v := range ch {
			fmt.Println(v)
		}

		i := 0
		for {
			i++
			if i > 10 {
				break
			}
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				fmt.Println("closed")
			}

		}
	}()

	close(ch)
	time.Sleep(1 * time.Second)
}
