package leetcode

func findJudge(N int, trust [][]int) int {
	isJudge := make([]int, N+1)
	// 遍历
	for _, v := range trust {
		// 信任别人 肯定不是法官 =-1后自然不可能是N-1
		isJudge[v[0]] = -1
		// 记录受到多少人信任 法官一定受到N-1人信任
		isJudge[v[1]]++
	}
	for i := 1; i <= N; i++ {
		if isJudge[i] == N-1 {
			return i
		}
	}
	return -1
}
