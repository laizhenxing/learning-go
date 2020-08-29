package int

import (
	"container/list"

	"strconv"
)

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	sL := len(s)
	mid := sL / 2
	for i := 0; i <= mid ; i++ {
		if s[i] != s[sL-i-1]{
			return false
		}
	}

	return true
}

// 1. 倒叙
// 2. 判断
func IsPalindrome1(x int) bool {
	s1 := list.List{}
	if x < 0 {
		return false
	}
	origin := x
	redirect := 0
	for x !=0 {
		redirect = redirect*10 + x%10
		x /= 10
	}
	return origin == redirect
}
