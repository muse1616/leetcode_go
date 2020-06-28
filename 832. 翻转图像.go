package leetcode

func flipAndInvertImage(A [][]int) [][]int {
	for i:=0;i<len(A);i++{
		//翻转
		for j:=0;j<len(A[i])/2;j++{
			A[i][j],A[i][len(A[i])-j-1] = A[i][len(A[i])-j-1],A[i][j]
		}
		for j:=0;j<len(A[i]);j++{
			A[i][j] ^= 1
		}
	}
	return A
}
