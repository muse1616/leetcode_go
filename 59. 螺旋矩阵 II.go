package leetcode

func generateMatrix(n int) [][]int {
	matrix:=make([][]int,n)
	for i:=0;i<len(matrix);i++{
		matrix[i] = make([]int,n)
	}
	count:=1
	//bound
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1

	for left <= right && top <= bottom {
		//1.向右
		for column := left; column <= right; column++ {
			 matrix[top][column] = count
			 count++
		}
		//2.向下
		for row := top + 1; row <= bottom; row++ {
			matrix[row][right] = count
			count++
		}
		// 此时需要判断是否剩下超过一列 一行 不然无需向左或向上
		if left < right && top < bottom {
			//3.向左
			for column := right - 1; column > left; column-- {
				matrix[bottom][column] = count
				count++
			}
			//4.向上
			for row := bottom; row > top; row-- {
				matrix[row][left] = count
				count++
			}
		}
		//边界收缩
		left++
		right--
		top++
		bottom--
	}

	return matrix
}
