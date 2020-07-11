package main

import "fmt"

func main() {
	var i int = 1
	i.PrintInt()
}

// cannot define new methods on non-local type int
func (i int) PrintInt()  {
	fmt.Println(i)
}