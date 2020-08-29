package array

import "testing"

func TestCombinationSum(t *testing.T) {
	nums1 := []int{2,3,6,7}
	target1 := 7
	expected1 := [][]int{
		{7},
		{2,2,3},
	}
	res1 := CombinationSum(nums1, target1)
	if len(res1) != len(expected1) {
		t.Fatalf("1. unexpected result: %v", res1)
	}

	nums2 := []int{2,3,5}
	target2 := 8
	expected2 := [][]int{
		{2,2,2,2},
		{2,3,3},
		{3,5},
	}
	res2 := CombinationSum(nums2, target2)
	if len(res2) != len(expected2) {
		t.Fatalf("1. unexpected result: %v", res2)
	}
}
