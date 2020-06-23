package leetcode

import "strconv"

func isValidSudoku(board [][]byte) bool {
	// 三个map 表示行列字表
	mapRow := make([]map[int]int, 9)
	mapCol := make([]map[int]int, 9)
	mapChild := make([]map[int]int, 9)
	// map中的map也需要初始化
	for i := 0; i < 9; i++ {
		mapRow[i] = make(map[int]int)
		mapCol[i] = make(map[int]int)
		mapChild[i] = make(map[int]int)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				// 转换为数字
				temp, _ := strconv.Atoi(string(board[i][j]))
				mapRow[i][temp]++
				mapCol[j][temp]++
				mapChild[(i/3)*3+j/3][temp]++
				if mapRow[i][temp] > 1 || mapCol[j][temp] > 1 || mapChild[(i/3)*3+j/3][temp] > 1 {
					return false
				}
			}

		}
	}

	return true

}
