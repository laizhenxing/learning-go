package main

import "fmt"

func f12(n int) (r int) {
	defer func() {
		r += n
		fmt.Println("recover r: ", r)
		recover()
	}()

	var ff func()

	defer ff()
	ff = func() {
		r += 2
		fmt.Println("ff r: ", r)
	}

	fmt.Println("r: ", r)
	fmt.Println("n: ", n)
	return n + 1
}

func main() {
	fmt.Println("test1: ", f12(3))
}