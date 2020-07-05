package leetcode

// dfs 搜索时将每个1标记为0 最后进行的dfs次数就是岛屿数量
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(&grid, i, j)
			}
		}
	}
	return ans
}

func dfs(grid *[][]byte, row int, column int) {
	//	重置为0
	(*grid)[row][column] = '0'
	//	深度遍历
	//  上 下 左 右
	if row-1 >= 0 && (*grid)[row-1][column] == '1' {
		dfs(grid, row-1, column)
	}
	if row+1 < len(*grid) && (*grid)[row+1][column] == '1' {
		dfs(grid, row+1, column)
	}
	if column-1 >= 0 && (*grid)[row][column-1] == '1' {
		dfs(grid, row, column-1)
	}
	if column+1 < len((*grid)[row]) && (*grid)[row][column+1] == '1' {
		dfs(grid, row, column+1)
	}
}
