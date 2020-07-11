package main

import "fmt"

func main() {
	// 自然数
	natrues := make(chan int)
	squares := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			natrues <- i
		}
		close(natrues)
	}()

	// 平方
	go func() {
		for natrue := range natrues {
			squares <- natrue * natrue
		}
		close(squares)
	}()
	// 打印
	for square := range squares {
		fmt.Println(square)
	}

}