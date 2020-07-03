package leetcode

// 反向操作 和边界的o相邻的o不需要被涂成x
func solve(board [][]byte) {
	//不需要变成x的点
	type point struct {
		x int //row
		y int //column
	}
	hash := make(map[point]bool, 10)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			// 值为X 或者不位于边界
			if board[i][j] == 'X' || (i != 0 && i != len(board)-1 && j != 0 && j != len(board[i])-1) {
				continue
			}
			// 已经在hash中的就不需要了
			//此处已经确定是位于边界的O
			p := point{i, j}
			if _, ok := hash[p]; ok {
				continue
			}
			// 相邻的O加入hash
			// 因为边界的o在外层循环时 所以不需要考虑同在边界的o
			// 使用队列保存
			queue := make([]point, 0)
			queue = append(queue, p)
			for len(queue) != 0 {
				tmpPoint := queue[0]
				if _, ok := hash[tmpPoint]; ok {
					//	出列
					queue = queue[1:]
					continue
				}
				hash[tmpPoint] = true
				// 上
				if tmpPoint.x-1 >= 0 {
					if board[tmpPoint.x-1][tmpPoint.y] == 'O' {
						p := point{tmpPoint.x - 1, tmpPoint.y}
						if _, ok := hash[p]; !ok {
							queue = append(queue, p)
						}
					}
				}
				// 下
				if tmpPoint.x+1 <= len(board)-1 {
					if board[tmpPoint.x+1][tmpPoint.y] == 'O' {
						p := point{tmpPoint.x + 1, tmpPoint.y}
						if _, ok := hash[p]; !ok {
							queue = append(queue, p)
						}
					}
				}
				// 右
				if tmpPoint.y+1 <= len(board[i])-1 {
					if board[tmpPoint.x][tmpPoint.y+1] == 'O' {
						p := point{tmpPoint.x, tmpPoint.y + 1}
						if _, ok := hash[p]; !ok {
							queue = append(queue, p)
						}
					}
				}
				// 左
				if tmpPoint.y-1 >= 0 {
					if board[tmpPoint.x][tmpPoint.y-1] == 'O' {
						p := point{tmpPoint.x, tmpPoint.y - 1}
						if _, ok := hash[p]; !ok {
							queue = append(queue, p)
						}
					}
				}
				//	出列
				queue = queue[1:]
			}
		}
	}
	//	将除了hash中的值全部变为X
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				p := point{i, j}
				if _, ok := hash[p]; !ok {
					board[i][j] = 'X'
				}
			}
		}
	}
}
