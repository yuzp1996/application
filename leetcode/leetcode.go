package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func deleteZero(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {

			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}
