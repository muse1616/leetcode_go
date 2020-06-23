package leetcode

import "math"

func thirdMax(nums []int) int {
	// 0 个 或 1 个
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[1]
	}

	max := nums[0]
	mid := math.MinInt32 - 1
	min := math.MinInt32 - 1

	for i := 1; i < len(nums); i++ {
		// 等于三数 跳过
		if nums[i] == max || nums[i] == mid || nums[i] == min {
			continue
		}

		if nums[i] > max {
			min, mid, max = mid, max, nums[i]
			continue
		}
		if nums[i] > mid {
			min, mid = mid, nums[i]
			continue
		}
		if nums[i] > min {
			min = nums[i]
			continue
		}
	}
	if min != math.MinInt32-1 {
		return min
	}

	return max
}
