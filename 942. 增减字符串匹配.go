package leetcode

func diStringMatch(S string) []int {
	result := make([]int, 0, len(S))
	// 记录最大值和最小值 避免每次的遍历
	minVal := 0
	maxVal := len(S)
	for i := 0; i < len(S); i++ {
		if S[i] == 'I' {
			result = append(result,minVal)
			minVal++
		} else {
			result = append(result,maxVal)
			maxVal--
		}

	}
	result =append(result,minVal)
	return result
}
