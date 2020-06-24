package leetcode

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	f := false
	//转为正数
	if num < 0 {
		num = -num
		f = true
	}
	var res string
	for num != 0 {
		//	头插
		//	如果商为0 则退出
		res = strconv.Itoa(num%7) + res
		num = num / 7
	}
	//加负号
	if f == true {
		res = "-" + res
	}
	return res

}
