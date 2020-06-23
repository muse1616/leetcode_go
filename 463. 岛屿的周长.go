package leetcode

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				var left,right,top,bottom int
				if j>=1{
					left = grid[i][j-1]
				}else {
					left=0
				}
				if j<=len(grid[0])-2{
					right = grid[i][j+1]
				}else {
					right=0
				}
				if i>=1{
					top = grid[i-1][j]
				}else {
					top = 0
				}
				if i<=len(grid)-2{
					bottom = grid[i+1][j]
				}else {
					bottom = 0
				}
				perimeter += 4 - left - right - top - bottom
			}
		}
	}

	return perimeter
}
