package array

import "testing"

func TestMaxArea(t *testing.T)  {
	height := []int{1,8,6,2,5,4,8,3,7}
	area := MaxArea(height)
	if area != 49 {
		t.Fatalf("unexpected result area: %d", area)
	}
}
