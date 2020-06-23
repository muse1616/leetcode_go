package leetcode

func findTheDifference(s string, t string) byte {
	m := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	for i := 0; i < len(s); i++ {
		m[s[i]]--
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 'a'
}
