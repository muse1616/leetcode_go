package leetcode

func countSegments(s string) int {

	num := 0
	flag := false
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if flag {
				num = num + 1
			}
			flag = false
		} else {
			flag = true
			continue
		}
	}
	if flag {
		num = num + 1
	}
	return num

}
