package leetcode

func exist(board [][]byte, word string) bool {

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if backtrace(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func backtrace(board [][]byte, word string, i, j, length int) bool {
	//单词长度找到
	if length == len(word) {
		return true
	}
	//board
	if i < 0 || j < 0 || i == len(board) || j == len(board[i]) {
		return false
	}
	if board[i][j] != word[length] {
		return false
	}

	//避免重复使用 回溯后重置
	tmp := board[i][j]
	board[i][j] = byte(' ')
	// 上
	if backtrace(board, word, i-1, j, length+1) {
		return true
	}
	// 右
	if backtrace(board, word, i+1, j, length+1) {
		return true
	}
	// 左
	if backtrace(board, word, i, j-1, length+1) {
		return true
	}
	// 下
	if backtrace(board, word, i, j+1, length+1) {
		return true
	}
	// 重置
	board[i][j] = tmp
	return false
}
