package leetcode

import "strconv"

var result bool

func isAdditiveNumber(num string) bool {
	if len(num) == 0 {
		return false
	}
	result = false
	//当前序列 当前数字起点 当前数字终点(当前字符) 所给字符串
	backtrace([]int{}, 0, 0, num)
	return result
}
func backtrace(arr []int, start int, end int, num string) {
	// 已经遍历到最后
	if end == len(num)-1 {
		currentNum, _ := strconv.Atoi(num[start : end+1])
		if len(arr) >= 2 && currentNum == arr[len(arr)-1]+arr[len(arr)-2] {
			result = true
			fmt.Println(arr)
		}
		return
	}
	// //第一位不能是0
	if num[start] == '0' {
		backtrace(append(arr, 0), start+1, start+1, num)
		return
	}
	//	直接放这个数字
	// 先验证放这个数字可不可以
	// 小于2 时放不放都可以
	if len(arr) < 2 {
		tmpNum, _ := strconv.Atoi(num[start : end+1])
		backtrace(append(arr, tmpNum), end+1, end+1, num)
		backtrace(arr, start, end+1, num)
	} else if len(arr) >= 2 {
		tmpNum, _ := strconv.Atoi(num[start : end+1])
		if tmpNum == arr[len(arr)-1]+arr[len(arr)-2] {
			backtrace(append(arr, tmpNum), end+1, end+1, num)
			backtrace(arr, start, end+1, num)
		} else {
			backtrace(arr, start, end+1, num)
		}
	}
}
