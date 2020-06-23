package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	d := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		d[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		d[t[i]]--
	}
	for _, v := range d {
		if v != 0 {
			return false
		}
	}
	return true
}
