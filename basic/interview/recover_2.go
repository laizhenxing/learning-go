package main

import "fmt"

func f()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover: %#v", err)
		}
	}()
	panic(1)
	panic(2)
}

func main() {
	f()
}
