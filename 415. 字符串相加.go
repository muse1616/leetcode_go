package leetcode

func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	var carry byte = 0
	var s []byte
	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			carry += num1[i] - '0'
			i--
		}
		if j >= 0 {
			carry += num2[j] - '0'
			j--
		}

		s = append(s, (carry%10)+'0')
		carry /= 10
	}
	n := len(s) - 1
	for i := 0; i < n; i++ {
		s[i], s[n] = s[n], s[i]
		n--
	}

	return string(s)

}
