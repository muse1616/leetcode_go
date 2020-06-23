package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	}

	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
		res = max(res, dp[i])
	}
	return res
}
