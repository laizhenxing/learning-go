package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func main() {
	names := []string{"a", "c", "b", "d", "F", "E", "i"}
	names1 := []string{"a", "c", "b", "d", "F", "E", "i"}
	sort.Sort(StringSlice(names))
	fmt.Println("main.StringSlice", names)
	sort.Strings(names1)
	fmt.Println("sort.Strings", names1)
	ints := []int{9,3,12,34,23,5,7,34,7,0,11,99}
	sort.Ints(ints)
	fmt.Println(ints)
}
