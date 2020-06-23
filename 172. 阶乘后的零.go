package leetcode

func trailingZeroes(n int) int {
	// 判断末尾的0 只需要看乘以10的个数 找10的因数 2和5 因为阶乘中 5 的数量 永远小于2的数量 所以看包含因数5的个数
	// 1*2*3*4（2*2）*5 故5永远少于2
	result:=0
	for n>0{
		n=n/5
		result+=n
	}
	return result
}