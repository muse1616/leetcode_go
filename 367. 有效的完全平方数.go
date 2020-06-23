package leetcode

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	left, right := 1, num/2
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else if mid*mid < num {
			left = mid + 1
		}

	}
	return false
}
