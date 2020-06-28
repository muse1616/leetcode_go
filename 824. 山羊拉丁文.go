package leetcode

import "strings"

func toGoatLatin(S string) string {
	result := ""
	strArr := strings.Split(S, " ")
	for index, word := range strArr {
		// 元音 a e i o u
		if word[0] == 'a' || word[0] == 'e' || word[0] == 'i' || word[0] == 'o' || word[0] == 'u' ||word[0] == 'A' || word[0] == 'E' || word[0] == 'I' || word[0] == 'O' || word[0] == 'U'{
			word += "ma"
		}else {
			b:=[]byte(word)
			b = append(b[1:],b[0])
			word = string(b)
			word+="ma"
		}
		//	加上相应数量的a
		for i := 0; i <= index; i++ {
			word += "a"
		}
		result += word + " "
	}
	return result[:len(result)-1]
}
