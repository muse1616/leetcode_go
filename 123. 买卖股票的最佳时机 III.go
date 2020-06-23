package leetcode

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	//		天数:len(prices) 当前是否持股 0,1 卖出的次数 0,1,2
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 3)
		}
	}
	//	初始条件
	//	第一天 未持股
	dp[0][0][0] = 0
	//	第一天 买入 赋值为负数即可
	dp[0][1][0] = -prices[0]
	//	第一天卖出的情况不存在
	dp[0][0][1] = -math.MaxInt16
	dp[0][0][2] = -math.MaxInt16
	dp[0][1][1] = -math.MaxInt16
	dp[0][1][2] = -math.MaxInt16

	//	状态转移
	for i := 1; i < len(dp); i++ {
		// 未持股，未卖出
		dp[i][0][0] = 0
		//未持股 卖出过一次 可能是今天卖的 也可能是之前卖的
		dp[i][0][1] = int(math.Max(float64(dp[i-1][1][0]+prices[i]), float64(dp[i-1][0][1])))
		//	未持股 卖出过两次 可能是今天卖的 也可能是之前卖的
		dp[i][0][2] = int(math.Max(float64(dp[i-1][1][1]+prices[i]), float64(dp[i-1][0][2])))
		//	持股 未卖出过 可能是今天买 也可能是之前买的
		dp[i][1][0] = int(math.Max(float64(dp[i-1][0][0]-prices[i]), float64(dp[i-1][1][0])))
		//	持股 卖出过 可能是今天买的 也可能是之前买的
		dp[i][1][1] = int(math.Max(float64(dp[i-1][0][1]-prices[i]), float64(dp[i-1][1][1])))
		//	持股 且卖出过两次 不可能
		dp[i][1][2] = -math.MaxInt16
	}
	//最后一天比较即可
	maxVal := int(math.Max(float64(dp[len(prices)-1][0][1]), float64(dp[len(prices)-1][0][2])))
	if maxVal < 0 {
		return 0
	}
	return maxVal
}
