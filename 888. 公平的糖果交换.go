package leetcode

func fairCandySwap(A []int, B []int) []int {
	sumA:=0
	sumB:=0
	for _,v:=range A{
		sumA+=v
	}
	for _,v:=range B{
		sumB+=v
	}
	gap:=(sumA-sumB)/2
	for i:=0;i<len(A);i++{
		for j:=0;j<len(B);j++{
			if A[i] - B[j] == gap{
				return []int{A[i],B[j]}
			}
		}
	}
	return nil
}
