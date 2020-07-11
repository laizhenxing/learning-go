package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 产生一个数
	go func() {
		for i := 1; i <= 10; i++ {
			ch1 <- i
		}
	}()

	// 计算一个数的平方
	go func() {
		for {
			r := <- ch1
			ch2 <- r * r
		}
	}()

	// 输出平方结果
	for {
		fmt.Println("result: ", <-ch2, time.Now().String())
	}
}
