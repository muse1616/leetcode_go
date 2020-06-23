package leetcode

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i+1)
	}
	//起始位置
	dp[0][0] = triangle[0][0]
	//dp填表
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			//最左边一列 只有一种走法
			if j == 0 {
				dp[i][j] = triangle[i][j] + dp[i-1][j]
			} else if j == i {
				// 最右侧一列
				dp[i][j] = triangle[i][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i-1][j-1]))) + triangle[i][j]
			}
		}
	}
	minPath := dp[n-1][0]
	for i := 0; i < len(dp[n-1]); i++ {
		if dp[n-1][i] < minPath {
			minPath = dp[n-1][i]
		}
	}
	return minPath

}
