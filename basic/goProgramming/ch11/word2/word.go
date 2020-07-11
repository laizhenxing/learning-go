package word

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	// 先分配好足够大的底层数组，避免在append调用是导致内存的多次重新分配，提升了50%的性能
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	n := len(letters)/2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
