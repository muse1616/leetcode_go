package leetcode



func imageSmoother(M [][]int) [][]int {
	rLen, cLen := len(M),len(M[0])
	avg := make([][]int,rLen)
	for i:=0;i<rLen;i++{
		avg[i] = make([]int,cLen)
		for j:=0;j<cLen;j++{
			sum,cnt := 0,0
			for m:=i-1;m<=i+1;m++{
				for n:=j-1;n<=j+1;n++{
					if m>=0 && m<rLen && n>=0 && n<cLen{
						sum += M[m][n]
						cnt++
					}
				}
			}
			avg[i][j] = sum/cnt
		}
	}
	return avg

}
