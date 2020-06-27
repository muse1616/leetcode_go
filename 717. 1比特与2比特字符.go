package leetcode

func isOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			i++
			continue
		}
		if i == len(bits)-1 && bits[i] == 0 {
			return true
		}
	}
	return false
}
