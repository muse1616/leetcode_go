package leetcode

import "math"

//dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
//dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
func maxProfit(prices []int) int {
	dp_i_0, dp_i_1 := 0, math.MinInt64
	// dp[i-2][0]
	dp_pre_0 := 0
	for i := 0; i < len(prices); i++ {
		tmp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_pre_0-prices[i], dp_i_1)
		dp_pre_0 = tmp
	}
	return dp_i_0

}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
