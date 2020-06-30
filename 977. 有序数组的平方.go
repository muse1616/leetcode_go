package leetcode

func sortedSquares(A []int) []int {
	ans:=make([]int,len(A))
	//	双指针
	left,right,index:=0,len(A)-1,len(A)-1

	for left<=right{
		if  A[left]*A[left]>A[right]*A[right]{
			ans[index] = A[left]*A[left]
			left++
		}else {
			ans[index] = A[right]*A[right]
			right--
		}
		index--
	}
	return ans
}
