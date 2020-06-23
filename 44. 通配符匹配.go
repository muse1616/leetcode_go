package leetcode

func isMatch(s string, p string) bool {
	//	dp初始化
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	//	边界条件
	dp[0][0] = true
	//第一行
	for i := 1; i <= n; i++ {
		if dp[0][i-1] == true && p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}

	//填表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i-1][j-1]
			} else if p[j-1] == '?' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}
