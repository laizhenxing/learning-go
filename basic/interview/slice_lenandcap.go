package main

import "fmt"

func main() {
	s := [3]int{1,2,3}
	a := s[:0]
	fmt.Printf("a len: %d, cap: %d\n", len(a), cap(a))
	b := s[:2]
	fmt.Printf("b len: %d, cap: %d\n", len(b), cap(b))
	c := s[1:2:cap(s)]
	fmt.Printf("c len: %d, cap: %d\n", len(c), cap(c))
}
