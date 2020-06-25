package leetcode

import "math"

func judgeSquareSum(c int) bool {
	min:=0
	max:=int(math.Sqrt(float64(c)))
	for min<=max{
		if min*min+max*max==c{
			return true
		}
		if min*min+max*max>c{
			max--
		}else {
			min++
		}
	}
	return false
}
