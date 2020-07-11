package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p := 1
	p = incr(&p)
	fmt.Println(p)

	s1 := add(1,2,3)
	fmt.Println(s1)
	sum := add([]int{1,2,34}...)
	fmt.Println(sum)
}

func add(args ...int) int {
	sum := 0
	for _, t := range args {
		sum += t
	}

	return sum
}