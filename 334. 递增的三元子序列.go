package leetcode

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums)<3{
		return false
	}
	small:=math.MaxInt64
	middle:=math.MaxInt64
	for _,v:=range nums{
		//前面两个都需要加 =
		if v<= small{
			small = v
		}else if v<=middle{
			middle = v
		}else if v>middle{
			return true
		}
	}
	return false
}
