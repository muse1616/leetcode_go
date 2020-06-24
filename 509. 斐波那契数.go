package leetcode

func fib(N int) int {
	return helper(N)
}
func helper(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return helper(n-1)+helper(n-2)
}
