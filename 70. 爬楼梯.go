package leetcode

func climbStairs(n int) int {
	m := make([]int, n+1)
	return a(0, n, m)
}

func a(i, n int, m []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if m[i] > 0 {
		return m[i]
	}
	m[i] = a(i+1, n, m) + a(i+2, n, m)
	return m[i]
}
