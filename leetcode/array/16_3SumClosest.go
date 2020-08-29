package array

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	nLen, res, diff := len(nums), 0, math.MaxInt32
	if nLen > 2 {
		sort.Ints(nums)
		for i := 0; i < nLen-2; i++ {
			for j, k := i+1, nLen-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
