package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	m := len(s1)
	n := len(s2)
	//	dp初始化
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[0][0] = true
			} else if i == 0 {
				//	第一行 j-1不会小于0
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				//	第一列
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				//	中间部分填图
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}

		}
	}
	return dp[m][n]
}
func main() {

	s1 := isInterleave("aabcc", "dbbca", "aadbbcbcac")
	fmt.Println(s1)
}
