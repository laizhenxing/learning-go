package main

import "fmt"

type S1 []int

func NewS1() S1 {
	return make(S1, 0)
}

func (s *S1) Add(elem int) *S1 {
	*s = append(*s, elem)
	fmt.Println(elem)
	return s
}

func main() {
	s := NewS1()
	defer s.Add(1).Add(2)
	s.Add(3)

	defer func() {
		s.Add(4).Add(5)
	}()
	s.Add(6)
}
