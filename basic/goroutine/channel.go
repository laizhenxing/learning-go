package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go add1(i, i+1, chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
	end := time.Now()
	t := end.Sub(start).Seconds()
	fmt.Println("执行耗时：", t)
}

func add1(a, b int, ch chan int)  {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	ch <- 1
}

