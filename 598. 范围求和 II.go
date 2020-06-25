package leetcode

import "math"

//只需要找ops里最小i和j的即可
func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	minI := 40000
	minJ := 40000
	for _, op := range ops {
		minI = int(math.Min(float64(minI), float64(op[0])))
		minJ = int(math.Min(float64(minJ), float64(op[1])))
	}
	return minI * minJ
}
