package leetcode

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	if len(s)==0{
		return 0
	}
	res :=0
	for i:=len(s)-1;i>=0;i--{
		if s[i]!=32{
			res++
		}else {
			return res
		}
	}

	return res
}
