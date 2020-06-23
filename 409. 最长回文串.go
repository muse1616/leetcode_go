package leetcode

func longestPalindrome(s string) int {
	result := 0
	hash := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hash[s[i]] ++
	}
	f := false
	for _, v := range hash {
		if v == 1 {
			f = true
			continue
		}
		// 偶数
		if v%2 == 0 {
			result += v
		} else {
			//	奇数
			result = result + v - 1
			f = true
		}
	}
	// 最大回文串为偶数 中间必能加一个
	if result%2 == 0 && f {
		result++
	}
	return result
}
