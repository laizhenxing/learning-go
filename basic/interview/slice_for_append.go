package main

import "fmt"

func main() {
	var a = []int{1,2,3}
	var r = make([]int, 0)

	for i, v := range a {
		if i == 0 {
			a = append(a,4,5,6)
		}
		r = append(r, v)
	}
	fmt.Println(r, a)
}
