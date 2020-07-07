package leetcode

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	//	初始条件
	dp[0] = 1

	//	指针 记录每个质数用了几次
	index := make([]int, len(primes))
	for i := 1; i < n; i++ {

		next := math.MaxInt64
		for k, v := range index {
			// dp[v]*primes[k] 指的是哪个质数和哪个dp结合
			next = min(next, primes[k]*dp[v])
		}
		for k, v := range index {
			if next == dp[v]*primes[k] {
				index[k]++
			}
		}
		dp[i] = next
	}

	return dp[len(dp)-1]
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
