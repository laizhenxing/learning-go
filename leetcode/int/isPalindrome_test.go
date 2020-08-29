package int

import "testing"

func TestIsPalindrome(t *testing.T) {
	x := -101
	if !IsPalindrome(x) {
		t.Fatalf("%d is not palindrome", x)
	}
}