package main

import "fmt"

func main() {
	defer_call()
}

// 发生panic,会先遍历该协程的defer,遇到recover,停止panic,返回recover继续执行
// 没有recover,遍历玩defer，再执行panic
// defer 先进后出
func defer_call()  {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("发生panic")
}
