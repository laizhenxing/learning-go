package main

import "fmt"

func main() {
	var ch chan int
	select {
	case v, ok := <-ch:
		fmt.Println(v, ok)
	default:
		fmt.Println("default")
	}
}
