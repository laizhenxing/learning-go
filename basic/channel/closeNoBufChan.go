package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)



	go func() {
		fmt.Println(<-ch1)
		fmt.Println("test1")
	}()

	go func() {
		ch1 <- 1
		fmt.Println("closed.")
		close(ch1)
		fmt.Println("ch1: ", <-ch1)
	}()

	time.Sleep(1*time.Second)
}
