package leetcode

import "math"

func titleToNumber(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result += int(math.Pow(float64(26), float64(len(s)-i-1)))*(int(s[i])-64)
	}
	return result
}
