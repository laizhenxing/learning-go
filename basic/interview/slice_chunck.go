package main

import "fmt"

func main() {
	chunck()
	compare()
}

func chunck()  {
	a := []int{1,2,3,4,5}
	t := a[3:4:4]	// 等价与 t = a[3:4]
	fmt.Println(t, cap(t))
}

func compare()  {
	// 数组是值类型，可比较
	a := [2]int{5,6}
	b := [2]int{5,6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}