package main

import "fmt"

func main() {
	m := make(map[string]int)
	// 不存在，但会返回一个零值
	v, ok := m["foo"]
	fmt.Println(ok, v)
	m["foo"]++
	fmt.Println(m["foo"])
}
