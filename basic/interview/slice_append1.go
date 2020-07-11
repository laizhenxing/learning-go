package main

import "fmt"

func main() {
	a := []int{7,8,9}
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)
}

func ap(a []int) {
	// append 会导致地城数组重新分配内存
	a = append(a, 10)
}

func app(a []int) {
	// 修改切片会修改底层数据内容
	a[0] = 1
}