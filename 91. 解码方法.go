package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	if s[n-1] == '0' {
		dp[n-1] = 0
	} else {
		dp[n-1] = 1
	}
	for i := n - 2; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
			continue
		}
		if (s[i]-'0')*10+(s[i+1]-'0') <= 26 {
			dp[i] = dp[i+1] + dp[i+2]
		} else {
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}
func main() {
	fmt.Println(numDecodings("1212"))
}
