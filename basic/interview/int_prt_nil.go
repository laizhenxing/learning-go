package main

import "fmt"

func Foo1(x interface{})  {
	fmt.Println("x: ", x)
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func main() {
	var i *int = nil
	if i == nil {
		fmt.Println(i)
	}
	Foo1(i)
}