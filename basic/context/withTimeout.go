package main

import (
	"context"
	"fmt"
	"time"
)

func runTimeout() {
	fmt.Println("timeout start...")
	timeoutHandler()
	fmt.Println("timeout stop...")
}

func timeoutHandler() {
	// 创建background子节点：withTimeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go doSth(ctx)

	// sleep
	time.Sleep(10 * time.Second)
	cancel()
}
