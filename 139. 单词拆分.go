package leetcode

func wordBreak(s string, wordDict []string) bool {
	//	初始化dp
	dp := make([]bool, len(s)+1)
	dp[0] = true

	//map判断是否在字典中
	d := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		d[v] = true
	}
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && d[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
