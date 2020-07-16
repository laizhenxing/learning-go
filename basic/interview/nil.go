package main

import "fmt"

func main() {
	nil := 123
	fmt.Println(nil)
	var _ map[int]string = nil	// cannot use nil (type int) as type map[int]string in assignment
}
