package leetcode

func findRepeatedDnaSequences(s string) []string {
	//结果数组
	var ans []string
	hash := make(map[string]bool, 10)

	// 记录每个字符串是否出现过
	for i := 0; i < len(s)-9; i++ {
		if _, ok := hash[s[i:i+10]]; ok {
			hash[s[i:i+10]] = true
		} else {
			hash[s[i:i+10]] = false
		}
	}
	for k, v := range hash {
		if v {
			ans = append(ans, k)
		}
	}
	return ans

}
