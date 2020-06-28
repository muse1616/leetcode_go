package leetcode

func rotateString(A string, B string) bool {
	if len(A) == 0 && len(B) == 0 {
		return true
	}
	b := []byte(A)
	for i := 0; i < len(b); i++ {
		//	讲第1位移到末尾
		b = append(b[1:], b[0])
		if string(b) == B {
			return true
		}
	}
	return false
}
