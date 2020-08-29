package practice

func UniqueSlice(s []int) []int {
	tmp := make([]int, 0)

	for _, v := range s {
		status := false
		for _, t := range tmp {
			if t == v {
				status = true
				break
			}
		}
		if !status {
			tmp = append(tmp, v)
		}
	}

	return tmp
}
