package main

import "fmt"

func main() {
	var a []int = nil
	a, a[0] = []int{1,2}, 9	// panic: runtime error: index out of range [0] with length 0
	fmt.Println(a)
}
