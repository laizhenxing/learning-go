// 时间复杂度：nlogn
// 不是原地排序，不稳定排序算法
package sort

import (
	"math/rand"
)

func Quick(s []int) []int {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}

	basic := s[rand.Intn(sLen)]

	left := make([]int, 0, sLen)
	middle := make([]int, 0, sLen)
	right := make([]int, 0, sLen)

	for _, v := range s {
		switch {
		case v < basic:
			left = append(left, v)
		case v == basic:
			middle = append(middle, v)
		case v > basic:
			right = append(right, v)

		}
	}

	left = Quick(left)
	right = Quick(right)

	left = append(left, middle...)
	left = append(left, right...)

	return left
}
