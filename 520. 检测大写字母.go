package leetcode

import (
	"strings"
)

func detectCapitalUse(word string) bool {
	//全小写或全大写
	if strings.ToLower(word) == word || strings.ToUpper(word) == word {
		return true
	}
	//首字母大写
	b := []byte(word)
	if len(b) == 1 || len(b) == 0 {
		return true
	}
	//首字母小写 直接不对
	if b[0] >= 97 && b[0] <= 122 {
		return false
	}
	for i := 1; i < len(b); i++ {
		//大写
		if b[i] < 97 {
			return false
		}
	}
	return true
}
