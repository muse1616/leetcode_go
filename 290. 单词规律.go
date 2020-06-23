package leetcode

func wordPattern(pattern string, str string) bool {
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)

	j := 0
	for i := 0; i < len(pattern); i++ {
		if j == len(str) {
			return false
		}
		word := ""
		for j < len(str) {
			if str[j] != ' ' {
				word += string(str[j])
				j++
			} else {
				j++
				break
			}
		}

		if _, ok := m1[pattern[i]]; ok == false {
			if _, ok := m2[word]; ok != false {
				return false
			}
			m1[pattern[i]] = word
			m2[word] = pattern[i]
			continue
		}
		if m1[pattern[i]] != word {
			return false
		}
	}
	if j != len(str) {
		return false
	}

	return true
}
