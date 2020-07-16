package main

import "fmt"

func F(n int) func() int  {
	return func() int {
		n++
		return n
	}
}

func main() {
	f := F(5)
	//fmt.Println(f())
	defer func(n int) {
		fmt.Println("defer func: ", f())	// 9
	}(f())	// 6

	defer fmt.Println("defer print: ", f())	// 7
	i := f()	// 8
	fmt.Println("print: ", i)
}
