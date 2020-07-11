package main

import "fmt"

func main() {
	var s1 []int
	//var s2 = []int{}
	if s1 == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	i := 48
	fmt.Println(string(i))
}
