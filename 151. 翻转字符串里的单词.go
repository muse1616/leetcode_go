package leetcode

import (
	"strings"
)

func reverseWords(s string) string {
	strArr := strings.Fields(s)
	if len(strArr) == 0 {
		return ""
	}
	i, j := 0, len(strArr)-1
	for i < j {
		strArr[i], strArr[j] = strArr[j], strArr[i]
		i++
		j--
	}
	ans := ""
	for _, v := range strArr {
		ans += v + " "
	}
	return ans[:len(ans)-1]
}
