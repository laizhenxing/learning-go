package main

import "fmt"

func main() {
	s1 := []string{"1", "a", "b", "c"}
	s2 := []string{"1", "a", "b", "c1"}
	e := isEqual(s1, s2)
	fmt.Println("s1==s2? ", e)
}

func isEqual(s,t []string) bool {
	if len(s) != len(t) {
		return false
	}
	for i, v := range s {
		if v != t[i] {
			return false
		}
	}

	return true
}
