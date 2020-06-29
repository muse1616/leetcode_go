package leetcode

import (
	"math"
	"strconv"
)

func binaryGap(N int) int {
	str := ""
	for ; N > 0; N /= 2 {
		str = strconv.Itoa(N%2) + str
	}
	maxDistance := 0
	for i := 0; i < len(str); i++ {
		if str[i] != '1' {
			continue
		}
		j := i + 1
		for ; j < len(str); j++ {
			if str[j] == '1' {
				break
			}
		}
		if j == len(str) {
			break
		}
		maxDistance = int(math.Max(float64(maxDistance), float64(j-i)))
	}
	return maxDistance
}
