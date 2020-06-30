package leetcode

func addToArrayForm(A []int, K int) []int {
	var i = len(A) - 1
	for ; i >= 0 && K > 0; i-- {
		// 保留进位 加到k上
		A[i], K = (A[i]+K)%10, (A[i]+K)/10
	}
	// K还有剩下
	for ; K != 0; K /= 10 {
		A = append([]int{K % 10}, A...)
	}
	return A
}
