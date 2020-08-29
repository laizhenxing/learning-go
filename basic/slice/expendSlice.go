package main

import "fmt"

func main() {
	s1 := make([]int, 3)
	fmt.Printf("len: %d, cap: %d\n", len(s1), cap(s1))
	s1 = append(s1,6,1,2,3,4,4,54,5,6,6,7,5,78,89,)
	fmt.Printf("len: %d, cap: %d\n", len(s1), cap(s1))
}
