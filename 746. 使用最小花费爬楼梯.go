package leetcode

import "math"

func minCostClimbingStairs(cost []int) int {
	a, b := 0, 0
	for i := 2; i <= len(cost); i++ {
		a, b = b, int(math.Min(float64(b+cost[i-1]), float64(a+cost[i-2])))
	}
	return b
}
