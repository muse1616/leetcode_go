package leetcode

func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	for i := 2; i <= n; i++ {
		if (1+i)*i/2 > n {
			return i - 1
		}
	}
	return 0
}
