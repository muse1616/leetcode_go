package leetcode

import "unicode"

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	i := 0
	j := len(s)-1
	for i < j {
		left:=rune(s[i])
		right:=rune(s[j])
		if isValid(left) == false && isValid(right) == false{
			i++
			j--
		}else if isValid(right)==false{
			j--
		} else if isValid(left) == false{
			i++
		} else if unicode.ToUpper(left) != unicode.ToUpper(right){
			return false
		}else {
			i++
			j--
		}
	}
	return true
}

func isValid(a rune) bool {
	return unicode.IsLetter(a) || unicode.IsDigit(a)
}
