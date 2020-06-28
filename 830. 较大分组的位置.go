package leetcode

func largeGroupPositions(S string) [][]int {
	b := []byte(S)
	result := make([][]int, 0, 1)
	for i := 0; i < len(b); {
		j := i + 1
		for ; j < len(b); {
			if b[j] == b[i] {
				j++
			} else {
				break
			}
		}
		if j-i >= 3 {
			result = append(result, []int{i, j - 1})
		}
		i = j

	}

	return result
}
