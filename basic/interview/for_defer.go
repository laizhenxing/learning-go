package main

import "fmt"

func main() {
	x := []int{1,2,3}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			fmt.Println(v)
		}()
		y[i] = &v
	}

	fmt.Println(*y[0], *y[1], *y[2])
}
