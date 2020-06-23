package leetcode

import (
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	// 去掉头尾空格
	str = strings.Trim(str, " ")
	// 去掉第一位正负号 从arr[1]开始算
	arr := []byte(str)
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		s, e := strconv.Atoi(string(arr[0]))
		if e == nil {
			return s
		}
		return 0
	}

	// 正负
	k := 0
	b := 0
	// 先判断第一位
	if arr[0] == 45 && (arr[1] >= 48 && arr[1] <= 57) {
		k = 1
	} else if arr[0] == 43 && (arr[1] >= 48 && arr[1] <= 57) {
		k = 1
		b = 1
	} else if (arr[0] < 48 || arr[0] > 57) && (arr[1] < 48 || arr[1] > 57) {
		return 0
	} else if arr[0] < 48 || arr[0] > 57 {
		return 0
	}

	res := []byte{}
	for i := k; i < len(arr); i++ {
		if arr[i] >= 48 && arr[i] <= 57 {
			res = append(res, arr[i])
		} else {
			if len(res) != 0 {
				break
			}
		}
	}
	number := 0
	if k == 0 || (k == 1 && b == 1) {
		// 正数
		number, _ = strconv.Atoi(string(res))
		if number > 1<<31-1 {
			return 1<<31 - 1
		}
	} else if k == 1 && b == 0 {
		// 负数
		number, _ = strconv.Atoi(string(res))
		number = 0 - number
		if number < -1<<31 {
			return -1 << 31
		}
	}

	return number
}
