package main

import "fmt"

func main() {
	s := make([]int, 5,5)
	s[0] = 1
	s[1] = 2
	change(s...)
	fmt.Println(s)
	change(s[0:2]...)
	fmt.Println(s)
}

func change(s ...int)  {
	s = append(s, 3)
	fmt.Println(s)
}
