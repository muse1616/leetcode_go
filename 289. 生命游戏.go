package leetcode

// 原来活->后来死 1->-2
// 原来活仍然活 1->1
// 原来死->后来活 0->-1
// 原来死仍然死 0->0
//全部更新完之后再遍历一次 -2改为0 -1改为1

func gameOfLife(board [][]int) {
	//遍历一次
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			judge(&board, i, j)
		}
	}
	//再遍历一次
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == -2 {
				board[i][j] = 0
			} else if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}
}

// 面板 行 列
// 原来是1 需要考虑四周的活细胞 即 1
// 原来是0 需要考虑四周的活细胞 即 1
// 所以四周的死细胞是无所谓的 边界周围默认为0即可
func judge(board *[][]int, row int, column int) {
	//	原来状态
	status := (*board)[row][column]
	//  0   1   2
	//  3  xx   4
	//  5   6   7
	c := make([]int, 8)
	if row > 0 && column > 0 {
		c[0] = (*board)[row-1][column-1]
	}
	if row > 0 {
		c[1] = (*board)[row-1][column]
	}
	if column < len((*board)[row])-1 && row > 0 {
		c[2] = (*board)[row-1][column+1]
	}
	if column > 0 {
		c[3] = (*board)[row][column-1]
	}
	if column < len((*board)[row])-1 {
		c[4] = (*board)[row][column+1]
	}
	if row < len(*board)-1 && column > 0 {
		c[5] = (*board)[row+1][column-1]
	}
	if row < len(*board)-1 {
		c[6] = (*board)[row+1][column]
	}
	if row < len(*board)-1 && column < len((*board)[row])-1 {
		c[7] = (*board)[row+1][column+1]
	}
	//	计算总和 此时 -2 需要看成1 -1看成0
	sum := 0
	for i := 0; i < len(c); i++ {
		if c[i] == -2 {
			c[i] = 1
		} else if c[i] == -1 {
			c[i] = 0
		}
		sum += c[i]
	}
	if sum < 2 && status == 1 {
		(*board)[row][column] = -2
		return
	} else if (sum == 2 || sum == 3) && status == 1 {
		return
	} else if sum > 3 && status == 1 {
		(*board)[row][column] = -2
		return
	} else if sum == 3 && status == 0 {
		(*board)[row][column] = -1
		return
	}
}
