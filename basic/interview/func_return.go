package main

import "fmt"

func getValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "exist", true
	}
	return nil, false	// cannot use nil as type string in return argument
}

func main() {
	intmap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	v, err := getValue(intmap, 3)
	fmt.Println(v, err)
}
