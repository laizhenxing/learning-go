package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	r := rate.NewLimiter(1, 5)
	ctx := context.Background()
	for {
		err :=  r.WaitN(ctx, 2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}
