package main

import (
	"fmt"
	"time"
)

func main() {
	taskch := make(chan int, 100)

	go worker1(taskch)

	// 塞任务
	for i := 0; i < 10; i++ {
		taskch <- i
	}

	// 等待1min
	select {
	case <-time.After(1 * time.Minute):
	}
}

func worker1(taskCh <-chan int) {
	// 启用5个工作协程
	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(1 * time.Second)
			}
		}(i)
	}
}
