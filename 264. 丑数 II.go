package leetcode

import "math"

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	//三个系数
	a2, a3, a5 := 0, 0, 0
	for i := 1; i < n; i++ {
		min := int(math.Min(math.Min(float64(dp[a3]*3), float64(dp[a5]*5)), float64(dp[a2]*2)))
		if min == dp[a2]*2 {
			a2++
		}
		if min == dp[a3]*3 {
			a3++
		}
		if min == dp[a5]*5 {
			a5++
		}
		dp[i] = min
	}

	return dp[len(dp)-1]
}
