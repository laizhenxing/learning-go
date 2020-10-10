package tools

// 找到连个字符串最长的公共前缀
func longestCommonPrefix(str1, str2 string) int {
	i := 0
	max := min(len(str1), len(str2))
	for i < max && str1[i] == str2[i] {
		i++
	}
	return i
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
