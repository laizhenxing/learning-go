package main

import (
	"fmt"

	"algorithm/sort"
)

func main() {
	s := []int{1,7,3,6,5,8,4,2,9,0}

	//sort.Bubble(s)
	s = sort.Merge(s)

	fmt.Println(s)
}
