package leetcode

//func isPowerOfTwo(n int) bool {
//	if n==1{
//		return true
//	}
//	return helper(n)
//}
//func helper(n int)bool  {
//	if n == 2{
//		return true
//	}
//	if n%2 != 0{
//		return false
//	}
//	if n>2{
//		return helper(n/2)
//	}
//	return false
//
//}
func isPowerOfTwo(n int) bool {
	//n如果是2的幂, 二进制下最高位为1, 其余全是0 n-1要么等于0, 要么二进制下都是1; n&(n-1)必为0
	return n > 0 && n&(n-1) == 0
}