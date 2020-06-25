package leetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
	//非合理行列
	if r*c != len(nums)*len(nums[0]) {
		return nums
	}
	arr:=make([]int,0,r*c)
	//结果矩阵初始化
	result := make([][]int, r)
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, c)
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			arr = append(arr, nums[i][j])
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			result[i][j] = arr[c*i+j]
		}
	}
	return result
}
