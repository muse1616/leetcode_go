package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	m1 := make(map[byte]int)
	for _, v := range magazine {
		m1[byte(v)]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if m1[ransomNote[i]] == 0 {
			return false
		}
		m1[ransomNote[i]] --
	}
	return true
}
