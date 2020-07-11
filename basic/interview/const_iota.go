package main

import "fmt"

const (
	a = iota
	b = iota
)
// iota 在 const 关键字出现时将被重置为0，const中每新增一行常量声明将使 iota 计数一次。
const (
	name = "name"
	age  = 12
	c    = iota
	d    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
