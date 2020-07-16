package main

import "fmt"

func main() {
	var s []int
	s = append(s, 1)
	fmt.Println(s)

	var m map[string]int
	m["test1"] = 1	// panic: assignment to entry in nil map
}

