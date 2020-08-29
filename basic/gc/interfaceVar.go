package main

import "fmt"

func main() {
	f := Fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci:%d\n", f())
	}
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
