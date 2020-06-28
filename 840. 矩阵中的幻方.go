package leetcode

func numMagicSquaresInside(grid [][]int) int {
	sum := 0
	for i := 0; i <= len(grid)-3; i++ {
		for j := 0; j <= len(grid[i])-3; j++ {
			if isMagic(grid, i, j) == true {
				sum++
			}
		}
	}
	return sum
}

// i j 作为左上角的矩形
func isMagic(grid [][]int, i int, j int) bool {

	// 9个数字个出现一次
	hash := make(map[int]int, 10)
	for m := i; m <= i+2; m++ {
		for n := j; n <= j+2; n++ {
			hash[grid[m][n]] ++
		}
	}
	//各出现一遍
	for m := 1; m <= 9; m++ {
		if hash[m] != 1 {
			return false
		}
	}

	sum := grid[i][j] + grid[i][j+1] + grid[i][j+2]
	//方向坐标
	x := []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2}
	y := []int{0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 2, 1, 0, 0, 0, 0, 1, 1, 1, 2, 2, 2}
	for m := 0; m < 8; m += 1 {
		tmpSum := 0
		for n := 0; n < 3; n++ {
			tmpSum += grid[i+x[m*3+n]][j+y[m*3+n]]
		}
		if tmpSum != sum {
			return false
		}
	}

	return true
}
