package leetcode

import "math"

func shortestToChar(S string, C byte) []int {
	var index []int
	for k, v := range S {
		if byte(v) == C {
			index = append(index, k)
		}
	}
	var result []int
	for k, _ := range S {
		min := len(S) + 1
		for i := 0; i < len(index); i++ {
			min = int(math.Min(float64(min), math.Abs(float64(k-index[i]))))
		}
		result =append(result,min)
	}

	return result
}
