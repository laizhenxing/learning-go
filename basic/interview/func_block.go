package main

import "fmt"

func main() {
	s := []int{1,2,3}
	x := 1
	fmt.Println(s[x])
	fmt.Println(x)	// 1

	{
		fmt.Println(x)	// 1
		i, x := 2, 3
		fmt.Println(i, x) // 2,3
	}

	fmt.Println(x)	// 1
}