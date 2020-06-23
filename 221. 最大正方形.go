package leetcode

import "math"

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m < 1 {
		return 0
	}
	n := len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	//最大边长
	maxSide := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				//左边 上边 左上方选一个 因为最小上方的被最大的包围了 画图理解
				dp[i][j] = 1 + int(math.Min(float64(dp[i-1][j-1]), math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))))
				maxSide = int(math.Max(float64(dp[i][j]), float64(maxSide)))
			}
		}
	}

	return maxSide * maxSide
}
