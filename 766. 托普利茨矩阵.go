package leetcode

func isToeplitzMatrix(matrix [][]int) bool {
	for row := 0; row < len(matrix)-1; row++ {
		for col := 0; col < len(matrix[row])-1; col++ {
			tmpRow := row
			tmpCol := col
			for tmpRow+1<len(matrix)&&tmpCol+1<len(matrix[0]){
				if matrix[tmpRow][tmpCol] != matrix[tmpRow+1][tmpCol+1]{
					return false
				}
				tmpRow++
				tmpCol++
			}
		}
	}
	return true
}
