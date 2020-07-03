package leetcode

var result [][]string

func partition(s string) [][]string {

	result = make([][]string, 0)
	// start和end表示当前字符串从序号start到end
	backtrace([]string{}, s, 0)
	return result
}
func backtrace(strArr []string, s string, start int) {
	// 回溯结束
	if start == len(s) {
		ans := make([]string, len(strArr))
		copy(ans, strArr)
		result = append(result, ans)
		return
	}
	//只有回文才加入strArr
	for i := start; i < len(s); i++ {
		if isPalindrome(s[start : i+1]) {
			backtrace(append(strArr, s[start:i+1]), s, i+1)
		}
	}

}

//判断是不是回文
func isPalindrome(str string) bool {
	if len(str) == 1 {
		return true
	}
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
