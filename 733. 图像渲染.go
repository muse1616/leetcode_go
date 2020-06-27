package leetcode

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[sr]) || image[sr][sc] == newColor {
		return image
	}
	//方向数组
	x := []int{-1, 0, 1, 0}
	y := []int{0, 1, 0, -1}
	oldColor := image[sr][sc]
	image[sr][sc] = newColor
	for i := 0; i <= 3; i++ {
		toSr := sr + x[i]
		toSc := sc + y[i]
		if toSr >= 0 && toSr < len(image) && toSc >= 0 && toSc < len(image[0]) && image[toSr][toSc] == oldColor {
			floodFill(image, toSr, toSc, newColor)
		}
	}
	return image
}
