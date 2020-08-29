package array

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2,7,11,15}
	target := 9
	expected := []int{0,1}
	res := TwoSum(nums, target)
	if len(res) != len(expected) {
		t.Fatalf("Len, unexpected result: %v", res)
	}
	if (res[0] != 0 && res[1] != 1) && (res[0] != 1 && res[1] != 0 ) {
		t.Fatalf("Res, unexpected result: %v", res)
	}
	fmt.Printf("expected result: %v\n", res)
}
