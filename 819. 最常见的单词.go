package leetcode

import (
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	// 使用非字母来分割字符串
	paragraph = strings.ToLower(paragraph)
	strArr :=strings.FieldsFunc(paragraph, func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	hash := make(map[string]int, 10)
	for _, v := range strArr {
		hash[v]++
	}
	// 去掉禁用
	for _, v := range banned {
		hash[v] = 0
	}
	maxV := 0
	for _, v := range hash {
		if v > maxV {
			maxV = v
		}
	}
	for k, v := range hash {
		if v == maxV {
			return k
		}
	}
	return ""
}
