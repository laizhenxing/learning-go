package main

import "fmt"

func t3() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		// 闭包延迟求值。
		funs = append(funs, func() {
			fmt.Println(&i, i)
		})
	}

	return funs
}

func main() {
	fns := t3()
	for _, f := range fns {
		f()
	}
}