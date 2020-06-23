package leetcode

func uniquePaths(m int, n int) int {
	//dp初始化
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	//	填表
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//边界条件 第一行第一列为1
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}