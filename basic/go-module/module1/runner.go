package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"go-module/module1/runner"
)

const timeout = 5 * time.Second

func main() {
	fmt.Println("开始运行...")

	// 初始化作业运行器
	r := runner.New(timeout)

	// 调度三个处理任务
	r.Add(createTask(), createTask(), createTask())

	// 启动作业运行器
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("作业程序因运行超时而终止")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("作业程序因系统发生中断事件而终止")
			os.Exit(2)
		}
	}
	log.Println("作业程序正常退出")
}

// 模拟后台处理任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

