package leetcode
func divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor==-1{
		return 2147483647
	}else{
		return dividend/divisor
	}
}