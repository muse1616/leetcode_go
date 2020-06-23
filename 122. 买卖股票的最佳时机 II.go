package leetcode

import "math"

func maxProfit(prices []int) int {
	n := len(prices)
	dp_i_0 := 0
	dp_i_1 := math.MinInt64
	for i := 0; i < n; i++ {
		tmp := dp_i_0
		dp_i_0 = int(math.Max(float64(dp_i_0), float64(dp_i_1+prices[i])))
		dp_i_1 = int(math.Max(float64(dp_i_1), float64(tmp-prices[i])))
	}
	return dp_i_0
}