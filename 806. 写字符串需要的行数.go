package leetcode

func numberOfLines(widths []int, S string) []int {
	var result []int
	sum := 0
	b := []byte(S)
	for i := 0; i < len(b); i++ {
		if 100-sum%100 >= widths[b[i]-'a'] {
			sum += widths[b[i]-'a']
		} else {
			sum = (sum/100+1)*100 + widths[b[i]-'a']
		}
	}
	if sum%100 == 0 {
		result = append(result, sum/100)
	} else {
		result = append(result, sum/100+1)
	}
	result = append(result, sum%100)

	return result
}
