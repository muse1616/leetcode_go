package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	// 空矩阵
	if len(matrix) == 0 {
		return false
	}
	//行 列
	row, col := len(matrix), len(matrix[0])
	//起始点
	i, j := row-1, 0
	for i >= 0 && j <= col-1 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			// 列数增加
			j++
		} else if matrix[i][j] > target {
			// 行数减少
			i--
		}
	}

	return false

}
