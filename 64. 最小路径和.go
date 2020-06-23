package leetcode

import "math"

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	//	初始化dp
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	//填表
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j >= 1 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 && i >= 1 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else if i >= 1 && j >= 1 {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}
