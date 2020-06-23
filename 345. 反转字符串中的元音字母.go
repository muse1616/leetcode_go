package leetcode

func reverseVowels(s string) string {
	byteArr := []byte(s)

	for left, right := 0, len(s)-1; left < right; {
		for !isVowel(s[left])&&left<right {
			left++
		}
		for !isVowel(s[right]) &&left<right{
			right--
		}
		if left >= right {
			break
		}
		byteArr[left], byteArr[right] = byteArr[right], byteArr[left]
		left++
		right--
	}
	return string(byteArr)
}

func isVowel(c byte) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return true
	}
	return false
}
