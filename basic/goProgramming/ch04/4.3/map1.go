package main

import "fmt"

func main() {
	var m1 map[string]int
	m1["test"] = 1
	fmt.Printf("%v\n", m1)
}
