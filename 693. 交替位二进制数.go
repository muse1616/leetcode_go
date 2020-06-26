package leetcode

func hasAlternatingBits(n int) bool {
	return toBin(n)
}
func toBin(n int) bool {
	flag := n % 2
	n = n / 2
	for n > 0 {
		if n%2 == flag {
			return false
		} else {
			flag = n % 2
			n = n / 2
		}
	}
	return true
}
