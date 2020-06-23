package leetcode

func countPrimes(n int) int {
	count := 0
	signs := make([]bool, n)
	for i := 2; i < n; i++ {
		// 非质数
		if signs[i] {
			continue
		}
		// 质数
		count++
		// 去除所有它的倍数 因为它的倍数一定不是素数
		for j := 2 * i; j < n ; j += i {
			signs[j] = true
		}
	}
	return count
}
