package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	runCancel()
	runTimeout()
}

func doSth(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		default:
			fmt.Printf("work %d sencod\n", i)
		}
		i++
	}
}


