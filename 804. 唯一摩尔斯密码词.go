package leetcode

func uniqueMorseRepresentations(words []string) int {
	morseStr := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	morseMap := map[string]int{}
	for _, v := range words {
		str := ""
		for _, v1 := range v {
			//-‘a’对应字母到morseStr的下标映射
			str += morseStr[v1-'a']
		}
		morseMap[str] = 1
	}
	return len(morseMap)

}
