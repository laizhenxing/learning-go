package main

import "fmt"

func main() {
	fmt.Println(f1()) // 1
	fmt.Println(f2()) // 5
	fmt.Println(f3()) // 1
}

func f1() (r int) {
	defer func() {
		r++
	}()

	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t += 5
	}()

	// 实际的执行顺序是：1. r = t 2. defer ... 3. return r
	return t
}

func f3() (r int) {
	defer func(r int) {
		r += 5
	}(r)

	// 实际的执行顺序是：1 r = 1; 2 defer ... return r
	return 1
}