package leetcode

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	//边界
	//第一列
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}
	//填表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//	若相同
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minOfThree(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}
func minOfThree(a int, b int, c int) int {
	if a < b && a < c {
		return a
	} else if b < c && b < a {
		return b
	}
	return c

}
