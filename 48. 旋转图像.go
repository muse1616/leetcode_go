package leetcode

func rotate(matrix [][]int) {
	//n := len(matrix) - 1
	//for i := 0; i < (len(matrix)+1)/2; i++ {
	//	for j := 0; j < len(matrix[i])/2; j++ {
	//		tmp := matrix[n-j][i]
	//		matrix[n-j][i] = matrix[n-i][n-j]
	//		matrix[n-i][n-j] = matrix[j][n-i]
	//		matrix[j][n-i] = matrix[i][j]
	//		matrix[i][j] = tmp
	//	}
	//}

//	转置+翻转
	for i:=0;i<len(matrix);i++{
		for j:=i;j<len(matrix);j++{
			matrix[i][j],matrix[j][i] = matrix[j][i],matrix[i][j]
		}
	}
//	翻转每一行
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix)/2;j++{
			matrix[i][j],matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j],matrix[i][j]
		}
	}
}
