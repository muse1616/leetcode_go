package leetcode

func hIndex(citations []int) int {
	length := len(citations)

	for i := range citations {
		if citations[i] >= length-i {
			return length - i
		}
	}
	return 0
}
