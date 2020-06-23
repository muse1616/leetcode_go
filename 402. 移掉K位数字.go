package leetcode

import "strings"

func removeKdigits(num string, k int) string {
	var result []byte
	for i := range num {
		for k > 0 && len(result) > 0 && result[len(result)-1] > num[i] { // 维持单调非递减
			result = result[:len(result)-1]
			k--
		}
		result = append(result, num[i])
	}
	//未删除满时 将末尾删除
	result = result[:len(result)-k]
	s := string(result)
	//去除开头0
	s = strings.TrimLeft(s, "0")
	if s == "" {
		return "0"
	}
	return s
}
