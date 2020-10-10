package main

import "fmt"

func main() {
	p := new([]int)
	fmt.Println(p)


	v := make([]int, 10)
	fmt.Println(v)

	(*p)[0] = 1
	fmt.Println(p)
	v[1] = 2
	fmt.Println(v)
}
