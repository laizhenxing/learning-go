package main

import "fmt"

type OneStruct struct {
	age int64
	name string
	email []string
	score map[string]float64
}

func main() {
	s1 := &OneStruct{
		age:   0,
		name:  "",
		email: nil,
		score: map[string]float64{"math": 100.0},
	}
	s2 := &OneStruct{
		age:   0,
		name:  "",
		email: nil,
		score: map[string]float64{"math": 100.0},
	}
	fmt.Println("s1 == s2?", s1 == s2)
}
