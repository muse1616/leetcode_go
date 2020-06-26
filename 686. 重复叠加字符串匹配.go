package leetcode

import "strings"

func repeatedStringMatch(A string, B string) int {
	c:=""
	for i:=1;i*len(A)<=2*len(A)+len(B);i++{
		c+=A
		if strings.Contains(c,B){
			return i
		}
	}
	return -1
}