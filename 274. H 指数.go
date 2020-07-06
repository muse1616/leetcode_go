package leetcode

import "sort"

func hIndex(citations []int) int {
	//for h := len(citations); h >= 1; h-- {
	//	count := 0
	//	for i := 0; i < len(citations); i++ {
	//		if citations[i] >= h {
	//			count++
	//			if count >= h {
	//				return h
	//			}
	//		}
	//	}
	//}
	//return 0
	sort.Ints(citations)
	length := len(citations)

	for i := range citations {
		if citations[i] >= length-i {
			return length - i
		}
	}
	return 0

}
