package leetcode

func isSubsequence(s string, t string) bool {
	start:=0

	for i:=0;i<len(t);i++ {
		if start == len(s){
			return  true
		}
		if s[start]!=t[i]{
			continue
		}
		start++
	}

	return start==len(s)
}
