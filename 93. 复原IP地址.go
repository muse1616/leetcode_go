package leetcode

import (
	"strconv"
)

var result []string

func restoreIpAddresses(s string) []string {
	// 0.0.0.0 ~ 255.255.255
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	result = make([]string, 0)
	//当次回溯ip 、 字符串 、数字个数 、字符串下标
	backtrace([]string{}, s, -1, len(s))
	return result
}

func backtrace(ipArr []string, s string, idx int, length int) {
	// 已到字符串结尾
	if idx == length-1 {
		// ip为s加三个点
		if len(ipArr) == 4 {
			result = append(result, ipArr[0]+"."+ipArr[1]+"."+ipArr[2]+"."+ipArr[3])
			return
		} else {
			return
		}
	}
	//下一位是0 就只能放0
	if string(s[idx+1]) == "0" {
		backtrace(append(ipArr, string(s[idx+1])), s, idx+1, length)
		return
	}
	//每次加一位数或二位数或三位数
	//	三位数需要判断 一位数两位数无需判断
	//	加一位数
	if idx+1 <= length-1 {
		backtrace(append(ipArr, string(s[idx+1])), s, idx+1, length)
	}
	if idx+2 <= length-1 {
		//	加两位数
		backtrace(append(ipArr, s[idx+1:idx+3]), s, idx+2, length)
	}
	//	加三位数需要判断
	if idx+3 <= length-1 {
		n, _ := strconv.Atoi(s[idx+1 : idx+4])
		if n <= 255 && n >= 0 {
			backtrace(append(ipArr, s[idx+1:idx+4]), s, idx+3, length)
		}
	}

}
