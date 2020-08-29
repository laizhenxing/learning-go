package array

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	nums1 := []int{1,2,0}
	exp1 := 3
	res1 := FirstMissingPositive(nums1)
	if res1 != exp1 {
		t.Fatalf("ex1. unexpected result: %v", res1)
	}

	nums2 := []int{3,4,-1,1}
	exp2 := 2
	res2 := FirstMissingPositive(nums2)
	if res2 != exp2 {
		t.Fatalf("ex2. unexpected result: %v", res2)
	}

	nums3 := []int{7,8,9,11,12}
	exp3 := 1
	res3 := FirstMissingPositive(nums3)
	if res3 != exp3 {
		t.Fatalf("ex3. unexpected result: %v", res3)
	}
}
