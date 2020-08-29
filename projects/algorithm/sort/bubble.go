// 时间复杂度：n^2
// 原地排序，稳定排序算法
package sort

func Bubble(s []int) []int {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(s)-1; i++ {
			if s[i+1] < s[i] {
				s[i], s[i+1] = s[i+1], s[i]
				swap = true
			}
		}
	}

	return s
}
