package leetcode

import "strconv"

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	res := ""
	flag := 0 //进位标识
	current := 0
	for i >= 0 || j >= 0 {
		Inta, Intb := 0, 0
		if i >= 0 {
			Inta = int(a[i] - '0')
		}
		if j >= 0 {
			Intb = int(b[j] - '0')
		}
		current = Inta + Intb + flag
		flag = 0
		if current >= 2 {
			flag = 1
			current = current - 2
		}
		cur := strconv.Itoa(current)
		res = cur + res
		i--
		j--
	}
	if flag == 1 {
		res = "1" + res
	}
	return res
}
