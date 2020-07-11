package main

import "fmt"

func main() {
	x := 2
	y := 7
	fmt.Println("x ^ y = ", x^y)

	fmt.Println("x&^y=", x&^y)
	fmt.Println("y&^x=", y&^x)
}
