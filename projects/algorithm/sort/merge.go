// 时间复杂度：nlogn
// 空间复杂度：n
// 不是原地排序，稳定排序算法
package sort

func Merge(s []int) []int {
	if len(s) < 2 {
		return s
	}

	mid := len(s) / 2
	var a = Merge(s[:mid])
	var b = Merge(s[mid:])

	return merge(a, b)
}

func merge(a, b []int) []int {
	fs := make([]int, 0, len(a)+len(b))
	i := 0 // ai
	j := 0 // bj

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			fs[i+j] = a[i]
			i++
		} else {
			fs[i+j] = b[j]
			j++
		}
	}

	// a[i:] 没有排完
	for i < len(a) {
		fs[i+j] = a[i]
		i++
	}
	// b[j:] 没有排完
	for j < len(b) {
		fs[i+j] = b[j]
		j++
	}

	return fs
}
