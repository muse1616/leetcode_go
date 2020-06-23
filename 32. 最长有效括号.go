package leetcode

import "math"

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	//	dp初始化
	dp := make([]int, len(s))
	dp[0] = 0
	max := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = 2
				if i-2 >= 0 {
					dp[i] = dp[i] + dp[i-2]
				}
			} else if dp[i-1] > 0 {
				if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
					if i-dp[i-1]-2 >= 0 {
						dp[i] = dp[i] + dp[i-dp[i-1]-2]
					}
				}
			}
			max = int(math.Max(float64(max), float64(dp[i])))
		}
	}

	return max
}
