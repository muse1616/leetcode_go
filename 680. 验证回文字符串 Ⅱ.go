package leetcode

func validPalindrome(s string) bool {

	b := []byte(s)
	left, right := 0, len(b)-1
	for left < right {
		if b[left] == b[right] {
			left++
			right--
		} else {
			flag1, flag2 := true, true
			for i, j := left+1, right; i < j; i, j = i+1, j-1 {
				if b[i] != b[j] {
					flag1 = false
					break
				}
			}

			for i, j := left, right-1; i < j; i, j = i+1, j-1 {
				if b[i] != b[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}
