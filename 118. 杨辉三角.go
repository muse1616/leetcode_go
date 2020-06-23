package leetcode

func generate(numRows int) [][]int {
	var result [][]int
	if numRows == 0 {
		return result
	}
	for i := 0; i < numRows; i++ {
		tmp := make([]int,0, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i ||i==0{
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, result[i-1][j-1]+result[i-1][j])
			}
		}
		result = append(result, tmp)
	}

	return result
}
