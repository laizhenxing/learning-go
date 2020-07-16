package main

import "fmt"

func main() {
	fmt.Println(dt1(1))	// 4
	fmt.Println(dt2(1))	// 3
}

func dt1(x int) (r int) {
	r = x
	defer func() {
		r += 3
	}()
	return r
}

func dt2(x int) (r int) {
	defer func() {
		r += x
	}()
	return 2
}
