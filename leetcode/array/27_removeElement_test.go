package array

import "testing"

func TestRemoveElement(t *testing.T) {
	nums1 := []int{3,2,2,3}
	val1 := 3
	expected1 := 2
	len1 := RemoveElement(nums1, val1)
	if len1 != expected1 {
		 t.Fatalf("1. unexpected result: %d", len1)
	}

	nums2 := []int{0,1,2,2,3,0,4,2}
	val2 := 2
	expected2 := 5
	len2 := RemoveElement(nums2, val2)
	if len2 != expected2 {
		t.Fatalf("2. unexpected result: %d", len2)
	}
}
