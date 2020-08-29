package array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums1 := []int{1,1,2}
	expected1 := 2
	len1 := RemoveDuplicates(nums1)
	if len1 != expected1 {
		t.Fatalf("unexpected result: %d", len1)
	}

	nums2 := []int{0,0,1,1,1,2,2,3,3,4}
	expected2 := 5
	len2 := RemoveDuplicates(nums2)
	if len2 != expected2 {
		t.Fatalf("unexpectd result: %d", len2)
	}
}
