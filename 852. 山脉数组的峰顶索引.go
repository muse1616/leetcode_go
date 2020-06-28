package leetcode

func peakIndexInMountainArray(A []int) int {
	for i:=0;i<len(A)-1;i++{
		if A[i]<A[i+1]{
			continue
		}else {
			return i
		}
	}
	return -1
}
