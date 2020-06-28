package leetcode

func transpose(A [][]int) [][]int {
//	A[i][j] <->A[j][i]
	result:=make([][]int,len(A[0]))
	for i:=0;i<len(result);i++{
		result[i] = make([]int,len(A))
	}
	for i:=0;i<len(A);i++{
		for j:=0;j<len(A[i]);j++{
			result[j][i] = A[i][j]
		}
	}
	return result
}
