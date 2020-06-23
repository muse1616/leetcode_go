package leetcode

func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if m1[s[i]] == 0 {
			if m2[t[i]] != 0 {
				return false
			}
			m1[s[i]] = t[i]
			m2[t[i]] = s[i]
			continue
		}
		if m1[s[i]] != t[i] {
			return false
		}
	}
	return true
}
