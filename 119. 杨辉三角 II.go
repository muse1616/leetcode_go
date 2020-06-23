package leetcode

func getRow(rowIndex int) []int {
	var result [][]int
	if rowIndex == 0 {
		return []int{}
	}
	for i := 0; i < rowIndex; i++ {
		tmp := make([]int,0, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i ||i==0{
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, result[i-1][j-1]+result[i-1][j])
			}
		}

		if i == rowIndex -1{
			return tmp
		}
		result = append(result, tmp)
	}
}