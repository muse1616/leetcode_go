package leetcode


// dp[i] i元钱需要多少硬币 即求dp[amount]
// 初始条件 dp[0]
// coins:c0 c1 c2 c3 c4...
// dp[i] = min(dp[i-c1]+1,dp[i-c2]+1......)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	// 初始条件
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[len(dp)-1] > amount{
		return -1
	}
	return dp[len(dp)-1]
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
