package leetcode

func reverseOnlyLetters(S string) string {
	b := []byte(S)
	left, right := 0, len(b)-1
	for left < right {
		for (b[left] < 65 || (b[left] >= 91 && b[left] <= 96)) && left<right{
			left++
		}
		for (b[right] < 65 || (b[right] >= 91 && b[right] <= 96)) && left<right{
			right--
		}
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}

	return string(b)
}
