package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	studs := []*Student{
		&Student{Name: "xiaoming"},
		&Student{Name: "xiaohong"},
		&Student{Name: "xiaolan"},
	}

	for i,v := range studs {
		fmt.Printf("%v, %v,\n", i, v)
	}
}
