package leetcode

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) || len(s1) == 0 {
		return false
	}
	//	dp初始化
	dp := make([][][]bool, len(s1))
	for i := 0; i < len(s1); i++ {
		dp[i] = make([][]bool, len(s2))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]bool, len(s1)+1)
		}
	}
	//切片l是1的时候判断两个字符相等
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1); j++ {
			dp[i][j][1] = s1[i] == s2[j]
		}
	}
	n := len(s1)
	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			for j := 0; j < n-l+1; j++ {
				//对长度为 l的S(i:i+L)和T(j:j+L)字符串切片 判断
				for k := 1; k < l; k++ {
					if dp[i][j][k] && dp[i+k][j+k][l-k] {
						dp[i][j][l] = true
						break
					}
					if dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}
				}

			}
		}
	}
	return dp[0][0][len(s1)]

}
