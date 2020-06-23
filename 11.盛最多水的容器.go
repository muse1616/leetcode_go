package leetcode

import "math"

func maxArea(height []int) int {
	max:=0
	for i:=0;i<len(height);i++{
		for j:=i;j<len(height);j++{
			if i!=j{
				min:=height[j]
				if height[i]<height[j]{
					min=height[i]
				}
				if min*int(math.Abs(float64(i-j)))>max{
					max = min*int(math.Abs(float64(i-j)))
				}
			}
		}
	}

	return max
}