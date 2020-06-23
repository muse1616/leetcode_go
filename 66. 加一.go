package leetcode

import "fmt"

func plusOne(digits []int) []int {
	l := len(digits)
	// 是否进位
	f := false
	for i := l - 1; i >= 0; i-- {
		add := 0
		if f == true || i == l-1 {
			add++
		}
		if digits[i]+add == 10 {
			f = true
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	fmt.Println(f)
	if f == true {
		res := make([]int, 0, len(digits)+1)
		res = append(res, 1)
		res = append(res, digits...)
		return res
	}
	return digits
}
