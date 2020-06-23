package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	j := 0
	i := len(matrix) - 1
	for i >= 0 && j <= len(matrix[0])-1 {
		if matrix[i][j] < target {
			j++
		} else if matrix[i][j] > target {
			i--
		} else {
			return true
		}
	}

	return false
}
