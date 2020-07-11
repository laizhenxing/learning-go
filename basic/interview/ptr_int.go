package main

import "fmt"

func incr1(i *int) int {
	*i++
	return *i
}

func main() {
	x := 1
	incr1(&x)
	fmt.Println(x)
}
