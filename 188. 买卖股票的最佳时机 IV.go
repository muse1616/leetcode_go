package leetcode

import "math"

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k < 1 {
		return 0
	}
	//k太大的时候 就没有约束了
	if k > len(prices)/2 {
		return maxProfit_k_inf(prices)
	}
	//		天数:len(prices) 当前是否持股 0,1 卖出的次数 0,1,2
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		//注意是k+1 因为k=0的情况用不到
		dp[i] = make([][]int, k+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 3)
		}
	}
	for i := 0; i < len(prices); i++ {
		for j := k; j >= 1; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[0]
				continue
			}
			dp[i][j][0] = int(math.Max(float64(dp[i-1][j][0]), float64(dp[i-1][j][1]+prices[i])))
			dp[i][j][1] = int(math.Max(float64(dp[i-1][j][1]), float64(dp[i-1][j-1][0]-prices[i])))
		}
	}

	return dp[len(prices)-1][k][0]
}

func maxProfit_k_inf(prices []int) int {

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
