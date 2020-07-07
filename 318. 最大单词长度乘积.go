package leetcode

import "strings"

func maxProduct(words []string) int {
	ans := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.ContainsAny(words[i], words[j]) {
				continue
			}
			ans = max(ans, len(words[i])*len(words[j]))
		}
	}
	return ans
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
