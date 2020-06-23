package leetcode

func isPowerOfFour(num int) bool {
	if num==1{
		return true
	}
	for num>4&&num%4==0{
		num /= 4
	}
	return num==4
}
