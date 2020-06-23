package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 1 {
			arr[i] = arr[i-1] + 1
		} else {
			arr[i] = 0
		}
	}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
