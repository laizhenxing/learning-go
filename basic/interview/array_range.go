package main

import "fmt"

func main() {
	var a = [5]int{1,2,3,4,5}
	var r [5]int

	// 参与循环的是 a 的副本,不是 a
	for i, v:=range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i]= v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
