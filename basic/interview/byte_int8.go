package main

import "fmt"

func main() {
	count := 0
	for i:= range [256]struct{}{} {
		m, n := byte(i), int8(i)
		if n == -n {
			count++
			fmt.Println("n: ", n)
		}
		if m == -m {
			count++
			fmt.Println("m: ", m, -m)
		}
	}
	fmt.Println(count)
}