package leetcode

import (
	"fmt"
	"strconv"
)
func hammingDistance(x int, y int) int {
	result := 0
	s1 := convertToBin(x)
	s2 := convertToBin(y)
	i, j := len(s1)-1, len(s2)-1
	for ; i >= 0 && j >= 0; {
		if s1[i] != s2[j] {
			result++
		}
		i--
		j--
	}
	if i!=-1{
		for k:=0;k<=i;k++{
			if s1[k]!='0'{
				result++
			}
		}
	}
	if j!=-1{
		for k:=0;k<=j;k++{
			if s2[k]!='0'{
				result++
			}
		}
	}
	return result

}

func convertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	for ; num > 0; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}
	return s
}
