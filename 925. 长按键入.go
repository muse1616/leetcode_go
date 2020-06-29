package leetcode

func isLongPressedName(name string, typed string) bool {
	if name[0] != typed[0] {
		return false
	}
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] == typed[j] {
			i++
			j++
		} else {
			if i == 0 {
				return false
			}
			if typed[j] == typed[j-1] {
				j++
			} else {
				return false
			}
		}
	}
	if i < len(name) || j < len(typed)-1{
		return false
	}
	return true
}
