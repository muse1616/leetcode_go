package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	sArr := make([]string, 0, len(nums))
	for _, v := range nums {
		sArr = append(sArr, strconv.Itoa(v))
	}
	sort.Slice(sArr, func(i, j int) bool {
		//保证高位数字一定大
		return sArr[i]+sArr[j] >= sArr[j]+sArr[i]
	})
	ans := strings.Join(sArr, "")
	if ans[0] == '0' {
		return "0"
	}
	return ans
}
