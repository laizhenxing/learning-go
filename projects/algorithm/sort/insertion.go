// 时间复杂度：n^2
// 原地排序，稳定排序算法
package sort

func Insertion(s []int) []int {
	for out := 1; out < len(s); out++ {
		tmp := s[out]
		in := out

		for ; in > 0 && s[in-1] >= tmp; in-- {
			s[in] = s[in-1]
		}
		s[in] = tmp
	}

	return s
}
