package leetcode

import "strconv"

func diffWaysToCompute(input string) []int {
	if isDigit(input) {
		tmp, _ := strconv.Atoi(input)
		return []int{tmp}
	}
	var ans []int
	for k, v := range input {
		//	如果是操作符 则计算两边的内容
		if v == '+' || v == '-' || v == '*' {
			left := diffWaysToCompute(input[:k])
			right := diffWaysToCompute(input[k+1:])
			// 左右结果排列组合即可
			for _, leftNum := range left {
				for _, rightNum := range right {
					var addNum int
					if v == '+' {
						addNum = leftNum + rightNum
					} else if v == '-' {
						addNum = leftNum - rightNum
					} else {
						addNum = leftNum * rightNum
					}
					ans = append(ans, addNum)
				}
			}
		}
	}
	return ans
}

//判断字符串中是否包含操作符
func isDigit(input string) bool {
	if _, err := strconv.Atoi(input); err != nil {
		return false
	}
	return true
}
