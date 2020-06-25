package leetcode

import "math"

func findMaxAverage(nums []int, k int) float64 {
	max:=math.MinInt32
	for i:=0;i<=len(nums)-k;i++{
		sum:=0
		for j:=0;j<k;j++{
			sum+=nums[i+j]
		}
		max = int(math.Max(float64(sum),float64(max)))
	}
	return float64(max)/float64(k)
}
