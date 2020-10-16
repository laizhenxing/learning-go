package main

import (
	"fmt"
	"math/rand"
)

func main() {
	chs := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3)	// 随机生成0-2的数字
	fmt.Printf("随机数：%d\n", index)
	chs[index] <- index

	select {
	case <-chs[0]:
		fmt.Println("chs[0]被选中")
	case <-chs[1]:
		fmt.Println("chs[1]被选中")
	case <-chs[2]:
		fmt.Println("chs[2]被选中")
	default:
		fmt.Println("没有chs被选中")
	}

}
