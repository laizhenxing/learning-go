package main

import "fmt"

func main() {
	s1 := make([]int, 4, 10)
	s1 = []int{1,2,3,4}
	s1 = appendInt(s1,5)
	s1 = append(s1, 6)
	fmt.Println(s1)

	ap()
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

func ap()  {
	s1 := []int{1,2,3,4}
	s2 := append(s1, 5)
	fmt.Printf("s1 = %v\n", s1)
	fmt.Printf("s2 = %v\n", s2)
}
