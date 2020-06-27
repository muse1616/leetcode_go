package leetcode

import "unicode"

func shortestCompletingWord(licensePlate string, words []string) string {
	var shortestWord string
	minLen := 16
	for _, v := range words {
		if len(v) < minLen {
			hash := make([]int, 26)
			for _, v2 := range v {
				hash[v2-'a']++
			}
			if isContain(hash, licensePlate) {
				minLen = len(v)
				shortestWord = v
			}
		}
	}
	return shortestWord
}
func isContain(list []int, word string) bool {
	for _, v := range word {
		if letter := v; unicode.IsLetter(letter) {
			if list[(letter|32)-'a'] <= 0 {
				return false
			}
			list[(letter|32)-'a']--
		}
	}
	return true
}
