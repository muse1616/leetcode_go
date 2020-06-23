package leetcode

func permute(nums []int) [][]int {
	var total [][]int
	backtrace(nums, 0, &total)
	return total
}
func backtrace(nums []int, idx int, total *[][]int) {
	if idx == len(nums)-1 {
		n := make([]int, len(nums))
		copy(n, nums)
		*total = append(*total, n)
		return
	}
	for i := idx; i < len(nums); i++ {
		nums[i], nums[idx] = nums[idx], nums[i]
		backtrace(nums, idx+1, total)
		nums[i], nums[idx] = nums[idx], nums[i]
	}
}
