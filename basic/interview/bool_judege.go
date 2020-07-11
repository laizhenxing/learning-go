package main

import "fmt"

func main() {
	x := interface{}(nil)
	y := (*int)(nil)
	a := y == x
	b := y == nil
	_, c := x.(interface{})
	fmt.Println(a, b, c)
}
