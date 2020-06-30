package leetcode

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var ans []int
	//bound
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1

	for left <= right && top <= bottom {
		//1.向右
		for column := left; column <= right; column++ {
			ans = append(ans, matrix[top][column])
		}
		//2.向下
		for row := top + 1; row <= bottom; row++ {
			ans = append(ans, matrix[row][right])
		}
		// 此时需要判断是否剩下超过一列 一行 不然无需向左或向上
		if left < right && top < bottom {
			//3.向左
			for column := right - 1; column > left; column-- {
				ans = append(ans, matrix[bottom][column])
			}
			//4.向上
			for row := bottom; row > top; row-- {
				ans = append(ans, matrix[row][left])
			}
		}
		//边界收缩
		left++
		right--
		top++
		bottom--
	}

	return ans
}
