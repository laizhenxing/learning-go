package main

import "fmt"

func t1(x int) (func(), func())  {
	return func() {
		fmt.Println(x)
		x += 10
	}, func() {
		// 闭包引用相同变量
		fmt.Println(x)
	}
}

func main() {
	a, b := t1(10)
	a()
	b()
}