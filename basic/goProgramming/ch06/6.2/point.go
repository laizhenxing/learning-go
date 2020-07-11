package main

import "fmt"

func main() {
	p := Point{
		a: 1,
		b: 2,
	}
	p.Sum()
}

type Point struct {
	a, b int
}

func (p Point) Sum() {
	fmt.Println("a+b=", p.a+p.b)
}
