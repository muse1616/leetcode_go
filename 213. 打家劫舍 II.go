package leetcode

import "math"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	//不包括nums[0]
	dp1 := make([]int, len(nums))
	//不包括nums[-1]
	dp2 := make([]int, len(nums))
	dp1[0], dp2[0] = 0, 0
	dp1[1], dp2[1] = nums[1], nums[0]
	// dp1 不包括nums[0]
	for i := 2; i < len(nums); i++ {
		dp1[i] = int(math.Max(float64(dp1[i-1]), float64(dp1[i-2]+nums[i])))
	}
	//dp2 不包括nums[-1]
	for i := 2; i < len(nums); i++ {
		dp2[i] = int(math.Max(float64(dp2[i-1]), float64(dp2[i-2]+nums[i-1])))
	}

	return int(math.Max(float64(dp1[len(dp1)-1]), float64(dp2[len(dp2)-1])))
}
