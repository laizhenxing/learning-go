package array

import (
	"fmt"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	nums1 := []int{10,1,2,7,6,1,5}
	target1 := 8
	expected1 := [][]int{
		{1,7},
		{1,2,5},
		{2,6},
		{1,1,6},
	}
	res1 := CombinationSum2(nums1, target1)
	if len(res1) != len(expected1) {
		t.Fatalf("1. unexpected result: %v", res1)
	}
	fmt.Println("res1: ", res1)

	nums2 := []int{2,5,2,1,2}
	target2 := 5
	expected2 := [][]int{
		{1,2,2},
		{5},
	}
	res2 := CombinationSum2(nums2, target2)
	if len(res2) != len(expected2) {
		t.Fatalf("1. unexpected result: %v", res2)
	}
	fmt.Println("res2: ", res2)
}
