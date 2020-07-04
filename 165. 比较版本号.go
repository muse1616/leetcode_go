package leetcode

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1Arr := strings.Split(version1, ".")
	v2Arr := strings.Split(version2, ".")
	i, j := 0, 0
	for i < len(v1Arr) && j < len(v2Arr) {
		n1, _ := strconv.Atoi(v1Arr[i])
		n2, _ := strconv.Atoi(v2Arr[j])
		if n1 > n2 {
			return 1
		} else if n1 < n2 {
			return -1
		}
		i++
		j++
	}
	//还有剩下
	for i < len(v1Arr) {
		n1, _ := strconv.Atoi(v1Arr[i])
		if n1 != 0 {
			return 1
		}
		i++
	}
	for j < len(v2Arr) {
		n1, _ := strconv.Atoi(v2Arr[j])
		if n1 != 0 {
			return -1
		}
		j++
	}
	return 0

}
