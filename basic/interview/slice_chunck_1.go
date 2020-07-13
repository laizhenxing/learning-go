package main

import "fmt"

func main() {
	x := make([]int, 2, 10)
	a := x[6:10]
	//b := x[6:]
	c := x[2:]
	fmt.Println(a, c)
}
