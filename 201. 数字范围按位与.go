package leetcode

func rangeBitwiseAnd(m int, n int) int {
	//ans :=m
	//for i:=m+1;i<=n;i++{
	//	if ans == 0{
	//		return 0
	//	}
	//	ans = ans & i
	//}
	//return ans

	for n > m {
		n = n & (n - 1)
	}
	return n
}
