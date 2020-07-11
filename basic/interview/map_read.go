package main

import "fmt"

func main() {
	readMap()
	readSlice()
}

type Person struct {
	name string
}

func readMap() {
	var m map[Person]int
	p := Person{
		name: "hello",
	}
	fmt.Println(m[p])
}

func changeSlice(num ...int) {
	// 传入的切片和函数内部的切片共用一个底层数组
	num[1] = 18
}

func readSlice()  {
	i := []int{5,6,7}
	changeSlice(i...)
	fmt.Println(i[1])
}

