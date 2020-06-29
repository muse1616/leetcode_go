package leetcode

import "math"

func isAlienSorted(words []string, order string) bool {
	//一个单词必然成立
	if len(words) == 1 {
		return true
	}
	//存储字母顺序表
	hash := make(map[byte]int)
	for k, v := range order {
		hash[byte(v)] = k
	}

	//遍历单词表
	for i := 1; i < len(words); i++ {
		length := int(math.Min(float64(len(words[i])),float64(len(words[i-1]))))
		if len(words[i]) < len(words[i-1]) {
			length = len(words[i])
		} else {
			length = len(words[i-1])
		}
		for j := 0; j < length; j++ {
			if hash[words[i][j]] > hash[words[i-1][j]] {
				break
			}
			if hash[words[i][j]] < hash[words[i-1][j]] {
				return false
			}
			if j == length-1 {
				if len(words[i]) < len(words[i-1]) {
					return false
				}
			}
		}
	}
	return true

}
