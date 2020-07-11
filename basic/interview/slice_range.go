package main

import "fmt"

func main() {
	var a = []int{1,2,3,4,5}
	var r [5]int

	// 参与循环的是 a 的副本,不是 a, 副本的指针依旧指向原slice的底层数组，对切片的修改都会反应到底层数组上
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
