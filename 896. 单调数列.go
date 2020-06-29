package leetcode

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}

	if A[len(A)-1] >= A[0] {
		//	递增
		for i := 0; i < len(A)-1; i++ {
			if A[i+1] < A[i] {
				return false
			}
		}
	} else {
		//	递减
		for i := 0; i < len(A)-1; i++ {
			if A[i+1] > A[i] {
				return false
			}
		}
	}
	return true
}
