package leetcode

import "strconv"

func isPalindrome(x int) bool {
	// 转为字符数组
	arr := []rune(strconv.Itoa(x))

	// 一个字符必是回文数
	if len(arr) == 1 {
		return true
	}

	// 偶数和基数
	if len(arr)%2 == 0 {
		for i := 0; i <= (len(arr)-1)/2; i++ {
			if arr[i] != arr[len(arr)-1-i] {
				return false
			}
		}
	} else {
		for i := 0; i <= (len(arr)-2)/2; i++ {
			if arr[i] != arr[len(arr)-1-i] {
				return false
			}
		}
	}

	return true
}
