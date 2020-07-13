package main

import "fmt"

const (
	x1 uint16 = 120
	y1
	s = "abc"
	t
)

func main() {
	fmt.Printf("%T %v\n", y1, y1)
	fmt.Printf("%T %v\n", t, t)

	const x2 = 124
	const y2 = 123
	fmt.Println(x2)
}