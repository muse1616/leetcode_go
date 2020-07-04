package leetcode

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		//递增数列
		if nums[left] <= nums[right] {
			break
		}
		middle := left + (right-left)/2
		if nums[middle] >= nums[left] {
			//	仍处于递增序列中
			left = middle + 1
		} else {
			right = middle
		}
	}
	return nums[left]
}
