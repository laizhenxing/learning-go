package main

import "fmt"

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowB() int {
	return w.i + 100
}

func (w Work) ShowA() int {
	return w.i + 10
}

func main() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowB())	// a.ShowB undefined (type A has no field or method ShowB)
	fmt.Println(b.ShowA())	// b.ShowA undefined (type A has no field or method ShowA)
}
