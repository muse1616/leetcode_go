package leetcode


func setZeroes(matrix [][]int) {
	row:=make(map[int]bool,10)
	col:=make(map[int]bool,10)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if row[i] == true{
				matrix[i][j] =0
			}
			if col[j] == true{
				matrix[i][j] = 0
			}
		}
	}
}
