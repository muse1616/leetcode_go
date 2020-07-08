package leetcode

func countBits(num int) []int {
	var ans []int
	for i:=0;i<=num;i++{
		ans = append(ans,toBinary(num))
	}
	return ans
}
func toBinary(num int)int{
	if num == 0{
		return 0
	}
	sum:=0
	for ;num>0;num/=2{
		if num%2 == 1{
			sum++
		}
	}
	return sum
}