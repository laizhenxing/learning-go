package array

import "sort"

func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)
	// 统计元素出现的次数
	counter := make(map[int]int)
	for _, v := range nums {
		counter[v]++
	}
	// 去重后的元素
	unique := make([]int, 0)
	for k := range counter {
		unique = append(unique, k)
	}
	// 排序
	sort.Ints(unique)

	for i := 0; i < len(unique); i++ {
		if (unique[i]*3 == 0) && counter[unique[i]] >= 3 {
			res = append(res, []int{unique[i], unique[i], unique[i]})
		}
		for j := i + 1; j < len(unique); j++ {
			if (unique[i]*2 + unique[j] == 0) && counter[unique[i]] > 1 {
				res = append(res, []int{unique[i], unique[i], unique[j]})
			}
			if (unique[j]*2 + unique[i] == 0) && counter[unique[j]] > 1 {
				res = append(res, []int{unique[i], unique[j], unique[j]})
			}
			c := 0 - unique[i] - unique[j]
			if c > unique[j] && counter[c] > 0 {
				res = append(res, []int{unique[i], unique[j], c})
			}
		}
	}

	return res
}
