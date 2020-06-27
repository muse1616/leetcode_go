package leetcode

import (
	"math"
)

func dominantIndex(nums []int) int {
	//if len(nums) == 1 {
	//	return 0
	//}
	//var cpy []int
	//for _, v := range nums {
	//	cpy = append(cpy, v)
	//}
	//sort.Ints(cpy)
	//if cpy[len(cpy)-1] >= cpy[len(cpy)-2]*2 {
	//	for i,v:=range nums{
	//		if v == cpy[len(cpy)-1]{
	//			return i
	//		}
	//	}
	//}
	//return -1
	max1, max2 := math.MinInt64, math.MinInt64
	maxindex := 0
	for i, num := range nums {
		if num > max1 {
			max2 = max1
			max1 = num
			maxindex = i
		} else if num > max2 {
			max2 = num
		}
	}
	if max1 < 2*max2 {
		return -1
	}

	return maxindex
}
