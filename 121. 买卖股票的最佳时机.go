package leetcode

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxVal := 0
	minVal := prices[0]
	for i := 0; i < len(prices); i++ {
		maxVal = int(math.Max(float64(maxVal), float64(prices[i]-minVal)))
		minVal = int(math.Min(float64(minVal), float64(prices[i])))
	}

	return maxVal
}
