package main

import "fmt"

func incrA() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer A: ", i)
	}()

	return i
}

func incrB() (r int) {
	defer func() {
		r++
		fmt.Println("defer B: ", r)
	}()
	return r
}

func main() {
	fmt.Println(incrA())
	fmt.Println(incrB())
}