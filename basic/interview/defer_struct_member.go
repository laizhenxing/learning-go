package main

import "fmt"

type person struct {
	age int
}

func main() {
	ps := &person{28}

	defer fmt.Println("1: ", ps.age)	// 28

	defer func(p *person) {
		fmt.Println("2: ", p.age)	// 29
	}(ps)

	defer func() {
		fmt.Println("3: ", ps.age)	// 29
	}()

	//ps.age = 29
	ps = &person{29}
}