package array

func FirstMissingPositive(nums []int) int {
	numsMap := make(map[int]int)
	for _, v := range nums {
		numsMap[v] = v
	}
	for i := 1; i < len(nums)+1; i++ {
		if _, ok := numsMap[i]; !ok {
			return i
		}
	}
	return len(nums) + 1
}
