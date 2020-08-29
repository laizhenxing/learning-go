package array

import "testing"

func TestThreeSumClosest(t *testing.T) {
	nums := []int{-1,2,1,-4}
	target := 1
	expected := 2
	res := ThreeSumClosest(nums, target)
	if res != expected {
		t.Fatalf("unexpected result: %d", res)
	}
}
