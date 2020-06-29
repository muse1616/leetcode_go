package leetcode

import "sort"

func smallestRangeI(A []int, K int) int {
	sort.Ints(A)
	if A[len(A)-1]-A[0]<=2*K{
		return 0
	}
	return A[len(A)-1]-A[0]-2*K
}
