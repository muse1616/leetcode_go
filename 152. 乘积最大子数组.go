package main

import "math"

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	val, maxVal, minVal := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		tmpMax, tmpMin := maxVal, minVal
		maxVal = int(math.Max(float64(tmpMax*nums[i]), math.Max(float64(nums[i]), float64(tmpMin*nums[i]))))
		minVal = int(math.Min(float64(tmpMin*nums[i]), math.Min(float64(nums[i]), float64(tmpMax*nums[i]))))
		val = int(math.Max(float64(maxVal), float64(val)))
	}

	return val
}
