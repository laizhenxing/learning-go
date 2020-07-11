package main

import "fmt"

func main() {
	s := "hello world!"
	t := s
	fmt.Println("before change: ", s)
	s = "yes! change it!"
	fmt.Println("s after change", s)
	fmt.Println("t after change", t)
}
