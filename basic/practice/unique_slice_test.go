package practice

import (
	"fmt"
	"testing"
)

func TestUniqueSlice(t *testing.T) {
	s := []int{1,1,2,2,3,3,4,4,5,5,6,6,7,7}
	ns := UniqueSlice(s)
	expect := []int{1,2,3,4,5,6,7}
	if len(ns) != len(expect) {
		t.Fatalf("unexpeted: %v", ns)
	}
	fmt.Println(expect)
}
