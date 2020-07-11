package main

import "fmt"

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{
		2,
		3,
	},
}

func main() {
	// map 的 value 是不可寻址的
	//m["foo"].x = 4	// cannot assign to struct field m["foo"].x in map
	fmt.Println(m["foo"].x)

	// 两种解决方法
	// 1. 使用临时变量
	// 2. 修改数据结构 map[string]*Math
}