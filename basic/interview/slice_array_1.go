package main

import "fmt"

func main() {
	var m map[int]bool // nil
	_ = m[123]
	var p *[5]string // nil
	for range p {
		fmt.Println(len(p))
	}
	var s []int // nil
	t := s[:]
	fmt.Println(t)
	s, s[0] = []int{1, 2}, 9	//  index out of range [0] with length 0
}
