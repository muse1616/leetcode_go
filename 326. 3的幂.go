package leetcode

//func isPowerOfThree(n int) bool {
//	return n > 0 && 1162261467 % n == 0
//}

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	for n > 3 && n%3 == 0 {
		n /= 3
	}
	return n == 3
}
