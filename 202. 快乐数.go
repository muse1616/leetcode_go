package leetcode
func isHappy(n int) bool {
	m:=make(map[int]bool)

	for n!=1{
		//出现过这个数字 说明循环
		if _,ok:=m[n];ok{
			return false
		}
		//计算下一位数字
		m[n] = true
		newNum:=0
		for n!=0{
			newNum += n%10 * (n%10)
			n /=10
		}
		n = newNum
	}
	return true
}