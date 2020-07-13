package main

import "fmt"

type T1 struct {
	x int
	y *int
}

func main() {
	i := 20
	t := T1{10, &i}

	p := &t.x

	*p++
	*p--

	t.y = p

	fmt.Println(t)
}
