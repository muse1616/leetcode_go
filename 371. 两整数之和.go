package leetcode

func getSum(a int, b int) int {
	for b != 0 {
		//异或相当于不考虑进位的加法
		sum := a ^ b
		//	与操作即每一位的进位
		carry := (a & b) << 1
		a = sum
		b = carry
	}
	return a
}
