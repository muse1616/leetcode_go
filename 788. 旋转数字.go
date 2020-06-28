package leetcode

func rotatedDigits(N int) int {
	sum:=0
	dp:=make([]int,N+1)
	copy(dp,[]int{0,0,1,-1,-1,1,1,-1,0,1})
	for i:=0;i<=N;i++{
		if dp[i/10] == -1 ||dp[i%10]==-1{
			dp[i] = -1
		}else if dp[i] = dp[i/10]|dp[i%10];dp[i]==1{
			sum++
		}
	}
	return sum


}
