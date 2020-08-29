// 时间复杂度：n^2
// 原地排序，不稳定排序算法
package sort

func Selection(s []int) []int {
	for i := 0; i < len(s); i++ {
		min := i
		for j := i+1; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}

	return s
}
