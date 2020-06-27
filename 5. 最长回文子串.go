package leetcode

func longestPalindrome(s string) string {
	//结果字符串
	var result string
	//二维dp初始化
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	//计数循环 counter表示有边界到左边界的距离
	for counter := 0; counter < len(s); counter++ {
		//左边界
		for i := 0; i+counter < len(s); i++ {
			// 右边界
			j := i + counter
			//边界条件
			if counter == 0 {
				//counter 即i==j 一个字母永远是回文 a
				dp[i][j] = 1
			} else if counter == 1 {
				// j = i+1 两个相同字母 aa
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				//状态转移方程
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 取大值 距离+1大于原答案 说明这个回文较长
			if dp[i][j] > 0 && j-i+1 > len(result) {
				result = s[i : i+counter+1]
			}
		}
	}
	return result
}
