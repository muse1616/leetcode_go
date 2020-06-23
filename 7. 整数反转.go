package leetcode

import "strconv"

func reverse(x int) int {
	flag := false
	if x < 0 {
		x = -x
		flag = true
	}
	arr := []rune(strconv.Itoa(x))
	// 反转
	var tmp int32
	if len(arr)%2 == 0 {
		for i := 0; i < (len(arr) / 2); i++ {
			tmp = arr[i]
			arr[i] = arr[len(arr)-i-1]
			arr[len(arr)-1-i] = tmp
		}
	}
	if len(arr)%2 != 0 {
		for i := 0; i <= (len(arr)-1)/2; i++ {
			tmp = arr[i]
			arr[i] = arr[len(arr)-i-1]
			arr[len(arr)-1-i] = tmp
		}
	}
	// arr转为数字
	res, _ := strconv.Atoi(string(arr))

	if res > 2147483647 || (0-res) < -2147483647 {
		return 0
	}

	if flag == false {
		return res
	}

	return 0 - res
}
