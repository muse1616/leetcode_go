package leetcode

import "math"

func surfaceArea(grid [][]int) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			sum += 4*grid[i][j] + 2
			if i-1 >= 0 {
				sum -= int(math.Min(float64(grid[i][j]), float64(grid[i-1][j])))
			}
			if j-1 >= 0 {
				sum -= int(math.Min(float64(grid[i][j]), float64(grid[i][j-1])))
			}
			if i+1 < len(grid) {
				sum -= int(math.Min(float64(grid[i][j]), float64(grid[i+1][j])))
			}
			if j+1 < len(grid) {
				sum -= int(math.Min(float64(grid[i][j]), float64(grid[i][j+1])))
			}
		}
	}

	return sum
}
