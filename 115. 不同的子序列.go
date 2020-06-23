package leetcode

func numDistinct(s string, t string) int {
	m := len(s)
	n := len(t)
	dp := make([][]int, n+1)
	//dp 初始哈
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	//	第一行初始化
	for i := 0; i <= m; i++ {
		dp[0][i] = 1
	}

	//	填表
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			//	字符相等
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[n][m]
}
