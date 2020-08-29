package main

import (
	"context"
	"fmt"
	"time"
)

func runCancel() {
	fmt.Println("cancel start...")
	cancelHandler()
	fmt.Println("cancel stop...")
}

func cancelHandler() {
	// 创建background
	ctx, cancel := context.WithCancel(context.Background())
	go doSth(ctx)

	// sleep
	time.Sleep(10 * time.Second)
	cancel()
}
