package leetcode

import (
	"math"
)

func countBinarySubstrings(s string) int {
	//	记录01连续出现的数组
	var record []int
	for i := 0; i < len(s); {
		//	内循环向后探索
		j := i + 1
		for ; j < len(s); j++ {
			if s[j] != s[j-1] {
				break
			}
		}
		record = append(record, j-i)
		i = j
	}
	//在record中每两个相邻数字取出较小的数字即可
	sum:=0
	for i:=0;i<len(record)-1;i++{
		sum+=int(math.Min(float64(record[i]),float64(record[i+1])))
	}
	return sum
}
