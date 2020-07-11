package main

import "fmt"

type info struct {
	r int
}

func work() (int, error) {
	return 1, nil
}

func main() {
	var i info

	// 不能使用短变量声明设置结构体字段值
	//i.r, err := work()	//  non-name i.r on left side of :=
	r, err := work()
	i.r = r
	fmt.Printf("info: %+v, %v\n", i, err)
}