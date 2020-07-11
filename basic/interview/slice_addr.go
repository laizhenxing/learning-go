package main

import "fmt"

func main() {
	s1 := []int{0,1,2,3}
	m := make(map[int]*int)

	// range循环中的k,v都是局部变量，内存地址并不会改变
	for k, v := range s1 {
		fmt.Println("&v=", &v)
		fmt.Println("&k=", &k)
		m[k] = &v
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
