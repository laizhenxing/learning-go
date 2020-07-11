package main

import "fmt"

var (
	size     = 1024
	max_size = size * 2
)

func main() {
	ap1()
	ap2()
}

func ap1() {
	list := new([]int)
	// *[]int 不能用于append()操作
	list = append(list, 1)
	fmt.Println(list)
}

func ap2() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
}
