package main

import "fmt"

func main() {
	// 不能给一个nil map赋值
	//var m map[string]int
	// 改进
	m := make(map[string]int)
	m["a"] = 1	// assignment to entry in nil map
	if v, ok := m["a"]; ok {
		fmt.Println(v)
	}
}
