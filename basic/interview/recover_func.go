package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		defer func() {
			fmt.Println(recover())
		}()
		panic(1)
	}()

	defer recover()
	panic(2)
}
