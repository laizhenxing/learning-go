package array

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[k] != nums[i] {
			k++
			nums[k] = nums[i]
		}
	}
	return k+1
}
