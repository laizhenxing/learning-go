package main

import "fmt"

func main() {
	var fn1 = func() {}
	var fn2 = func() {}

	// 函数只能和nil做比较
	if fn1 != fn2 {	//  invalid operation: fn1 != fn2 (func can only be compared to nil)
		fmt.Println("fn1 not equal fn2")
	}
}
