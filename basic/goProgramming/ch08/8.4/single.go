package main

import "fmt"

func main() {
	out := make(chan int)
	in := make(chan int)
	// 发送chan
	go counter(in)
	// 平方
	go squarer(out, in)
	// 打印
	print(out)
}

// 将输入送入chan
// close只能用于发送chan
func counter(out chan<- int) {
	for i := 1; i < 100; i++ {
		out <- i
	}
	close(out)
}

// 平方
// out 用于存储结果，将结果发送到out chan
// in 用于原始数据输入
func squarer(out chan<- int, in <-chan int) {
	for n := range in {
		out <- n * n
	}
	close(out)
}

// print
func print(in <-chan int) {
	for i := range in{
		fmt.Println(i)
	}
}
