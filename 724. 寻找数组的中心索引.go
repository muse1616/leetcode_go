package leetcode

func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		left := 0
		right := 0
		for j := 0; j < i; j++ {
			left += nums[j]
		}
		for j := i + 1; j < len(nums); j++ {
			right += nums[j]
		}
		if left == right {
			return i
		}
	}
	return -1
}
