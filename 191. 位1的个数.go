package leetcode

import "strconv"

func hammingWeight(num uint32) int {
	var count int
	for {
		if num == 0 {
			break
		}
		if num % 2 == 1 {
			count++
		}
		num = num / 2
	}
	return count
}