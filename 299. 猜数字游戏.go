package leetcode

import (
	"strconv"
)

func getHint(secret string, guess string) string {
	bull := 0
	cow := 0
	m := make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		m[secret[i]]++
	}
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bull++
			m[secret[i]]--
		}
	}
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			continue
		}
		if m[guess[i]] > 0 {
			cow++
			m[guess[i]]--
		}
	}
	return strconv.Itoa(bull) + "A" + strconv.Itoa(cow) + "B"
}
