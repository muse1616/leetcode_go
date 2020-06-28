package leetcode

func letterCasePermutation(S string) []string {
	var result []string
	b := []byte(S)
	for k, c := range b {
		if k == 0 {
			if c >= 48 && c <= 57 {
				result = append(result, string(c))
			} else {
				result = append(result, string(c))
				if c >= 65 && c <= 90 {
					result = append(result, string(c+32))
				} else {
					result = append(result, string(c-32))
				}
			}
			continue
		}
		//	如果是数字 则在result中所有字符串后添加一个数字即可
		if c >= 48 && c <= 57 {
			for i := 0; i < len(result); i++ {
				result[i] += string(c)
			}
		} else {
			//	如果是字母 则复制一份result 在第一份添加本身 第二份添加变换
			result = append(result, result...)
			for i := 0; i < len(result)/2; i++ {
				result[i] += string(c)
			}
			for i := len(result) / 2; i < len(result); i++ {
				if c >= 65 && c <= 90 {
					result[i] += string(c + 32)
				} else {
					result[i] += string(c - 32)
				}
			}
		}
	}

	return result
}
