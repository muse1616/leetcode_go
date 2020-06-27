package leetcode

import "strconv"

func selfDividingNumbers(left int, right int) []int {
	var result []int
	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) == true {
			result = append(result, i)
		}
	}
	return result
}

//判断是否是自除数
func isSelfDividingNumber(number int) bool {
	b := []byte(strconv.Itoa(number))
	for i := 0; i < len(b); i++ {
		a, _ := strconv.Atoi(string(b[i]))
		if b[i] == '0' || number%a != 0 {
			return false
		}
	}
	return true
}
