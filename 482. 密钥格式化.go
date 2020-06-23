package leetcode

import "strings"

func licenseKeyFormatting(S string, K int) string {
	//去掉所有-
	S = strings.ReplaceAll(S, "-", "")
	if len(S) == 0 {
		return ""
	}
	//第一组的字符数
	groupOne := len(S) % K
	if groupOne == 0 {
		groupOne = K
	}
	var result []byte
	b := []byte(S)
	//第一部分
	result = append(result, b[:groupOne]...)
	if groupOne == len(b) {
		return strings.ToUpper(string(result))
	}
	result = append(result, '-')

	for i := groupOne; i < len(b); i++ {
		if (i-groupOne)%K == 0 && i != groupOne {
			result = append(result, '-')
			result = append(result, b[i])
		} else {
			result = append(result, b[i])
		}

	}
	return strings.ToUpper(string(result))
}
