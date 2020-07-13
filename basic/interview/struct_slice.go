package main

import "fmt"

type T struct {
	ls []int
}

func main() {
	t := T{ls: []int{1,2,3}}

	foo1(t)
	fmt.Println(t.ls[0])
}

func foo1(t T)  {
	t.ls[0] = 100
}