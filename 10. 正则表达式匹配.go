package leetcode

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	//	dp初始化 需要大一圈以放置边界条件
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	// 边界条件
	//	第一列
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		dp[i][0] = false
	}
	//	第一行
	for i := 2; i <= n; i++ {
		if p[i-1] != '*' {
			dp[0][i] = false
		} else {
			dp[0][i] = dp[0][i-2]
		}
	}

	//填表
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if j >= 2 && p[j-1] == '*' {
				t1 := dp[i-1][j] && (p[j-2] == s[i-1] || p[j-2] == '.')
				t2 := dp[i][j-2]
				dp[i][j] = t1 || t2
			} else if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = false
			}
		}
	}

	//右下角即为结果
	return dp[m][n]
}
