package leetcode

func validMountainArray(A []int) bool {

	if len(A) < 3 {
		return false
	}
	//递增标志
	isUp := false
	if A[0] < A[1] {
		isUp = true
	} else {
		return false
	}

	for i := 1; i< len(A)-1; i++ {
		if isUp {
			if A[i+1] == A[i] {
				return false
			}
			if A[i+1] < A[i] {
				isUp = false
				continue
			}
			continue
		}
		if A[i+1] >= A[i] {
			return false
		}
	}
	if isUp == true {
		return false
	}
	return true

}
