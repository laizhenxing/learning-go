package main

import "fmt"

func main() {
	var x interface{}
	var y interface{} = []int{1,2,3}
	fmt.Println(x == x)	// true
	fmt.Println(x == y)	// false
	_ = y == y	// panic: runtime error: comparing uncomparable type []int
}
