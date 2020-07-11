package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

func main() {
	var (
		expr *cronexpr.Expression
		err error
	)

	// 每分钟执行一次
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	// 当前时间
	now := time.Now()
	// 下次调用时间
	nextTime := expr.Next(now)
	fmt.Println(now, nextTime)

	// 等待这个定时器超时
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了", nextTime)
	})

	time.Sleep(5 * time.Second)
}