package leetcode

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	for {
		if num == 1 {
			return true
		}
		if num%2 == 0 {
			num /= 2
			continue
		}
		if num%3 == 0 {
			num /= 3
			continue
		}
		if num%5 == 0 {
			num /= 5
			continue
		}
		return false
	}
}
