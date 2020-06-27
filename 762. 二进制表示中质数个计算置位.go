package leetcode

func countPrimeSetBits(L int, R int) int {
	result := 0
	for i := L; i <= R; i++ {
		result += helper(i)
	}
	return result
}
func helper(num int) int {
	count := 0
	//	转换成二级制 并计算1的个数
	for ; num > 0; num /= 2 {
		if num%2 == 1 {
			count++
		}
	}
	if count == 1 {
		return 0
	}
	if count == 2 {
		return 1
	}
	//	判断是不是质数
	for i := 2; i < count; i++ {
		if count%i == 0 {
			return 0
		}
	}
	//	返回1或0
	return 1
}
