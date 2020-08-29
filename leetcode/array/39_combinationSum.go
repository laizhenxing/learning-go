package array

import "sort"

// 元素可重复使用
func CombinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findCombinationSum(candidates, target, 0, c, &res)
	return res
}

func findCombinationSum(nums []int, target, index int, c []int, res *[][]int)  {
	if target <= 0 {
		if target == 0 {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}
	for i := index; i < len(nums); i++ {
		if nums[i] > target {	// 剪枝优化，如果 nums[i] 不满足目标，则比nums[i]大的数也不满足
			break
		}
		c = append(c, nums[i])
		findCombinationSum(nums, target - nums[i], i, c, res)
		c = c[:len(c)-1]
	}
}
