package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(1 * time.Second)
	// 每隔1s执行一次定时任务
	for {
		select {
		case <-ticker:
			fmt.Println("time ticker")
		}
	}

}
