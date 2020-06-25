package leetcode

import "math"

func findLHS(nums []int) int {
	hash := make(map[int]int)
	for _, v := range nums {
		if _, ok := hash[v]; !ok {
			hash[v] = 1
		} else {
			hash[v]++
		}
	}
	max := 0
	for k, v := range hash {
		//i i-1
		if value, ok := hash[k-1]; ok {
			max = int(math.Max(float64(value+v), float64(max)))
		}
		//i i+1
		if value, ok := hash[k+1]; ok {
			max = int(math.Max(float64(value+v), float64(max)))
		}
	}
	return max

}
