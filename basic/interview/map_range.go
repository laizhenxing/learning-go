package main

import "fmt"

func main() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	// map是无序的，每次循环开始的顺序不一样
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is: ", counter)
}
