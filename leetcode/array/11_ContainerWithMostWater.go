package array

func MaxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1

	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		area := width * high
		if area > max {
			max = area
		}
	}

	return max
}
