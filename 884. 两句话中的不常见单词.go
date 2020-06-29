package leetcode

import (
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	//分割字符串 防止空格数量不定
	arr1 := strings.FieldsFunc(A, func(r rune) bool {
		if r == ' ' {
			return true
		}
		return false
	})
	arr2 := strings.FieldsFunc(B, func(r rune) bool {
		if r == ' ' {
			return true
		}
		return false
	})
	//字母出现表
	hash1 := make(map[string]int, 10)
	hash2 := make(map[string]int, 10)
	var result []string

	for i := 0; i < len(arr1); i++ {
		hash1[arr1[i]]++
	}
	for i := 0; i < len(arr2); i++ {
		hash2[arr2[i]]++
	}
	//  hash1 中出现一次 hash2 不出现   hash2中出现一次 hash1不出现
	for k,v:=range hash1{
		if _,ok:=hash2[k];!ok&&v==1{
			result =  append(result,k)
		}
	}
	for k,v:=range hash2{
		if _,ok:=hash1[k];!ok&&v==1{
			result =  append(result,k)
		}
	}

	return result
}
