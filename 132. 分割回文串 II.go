package leetcode

import "math"

func minCut(s string) int {
	//无需分割
	if len(s) <= 1 {
		return 0
	}
	//dp初始化
	dp := make([]int, len(s))

	//默认值是完全切割
	for i := 0; i < len(s); i++ {
		dp[i] = i
	}
	for i := 1; i < len(s); i++ {
		//	先判定0--i是不是回文 是回文这一项无需切割
		if check(s, 0, i) == true {
			dp[i] = 0
		} else {
			//	不是回文 s[j+1:i]依次遍历 若s[j+1:i]是回文 相当于在dp[j]的基础上多了一次切割
			for j := 0; j < i; j++ {
				if check(s, j+1, i) == true {
					dp[i] = int(math.Min(float64(dp[i]), float64(dp[j])+1))
				}
			}
		}
	}

	//最后一项dp是总切割数目
	return dp[len(s)-1]
}

//判断是不是回文
func check(s string, left, right int) bool {
	for ; left < right; {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
