package array

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := ThreeSum(nums)
	if len(res) != 2 {
		t.Fatalf("unexpected result: %v", res)
	}

	fmt.Printf("expected result: %v\n", res)
}
