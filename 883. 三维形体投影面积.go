package leetcode

import "math"

func projectionArea(grid [][]int) int {
	//	xy面积是中长度 （!=0）
	//Sxy := 0
	//	xz面积 x相同时 取数值大的 加起来
	// 比较grid[i]中最大的即可
	//Sxz := 0
	//  yz面积 y相同时 取数值大的 加起来
	//Syz := 0
	s:=0
	for i := 0; i < len(grid); i++ {
		m1 := 0
		m2 := 0
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 {
				s++
			}
			m1 = int(math.Max(float64(m1), float64(grid[i][j])))
			m2 = int(math.Max(float64(m2), float64(grid[j][i])))
		}
		s += m1
		s += m2
	}
	return s
}
