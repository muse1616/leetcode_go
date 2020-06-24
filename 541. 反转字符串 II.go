package leetcode

func reverseStr(s string, k int) string {

	b := []byte(s)

	for i := 0; i < len(b); i += 2 * k {
		left := i
		right := i + k - 1
		if right >= len(b) {
			right = len(b) - 1
		}
		for left < right && left < len(b) {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
	}
	return string(b)

}
