package leetcode

var ans [][]int

func combinationSum3(k int, n int) [][]int {
	ans = make([][]int, 0)
	backtrace(k, n, []int{}, 1)
	return ans
}

func backtrace(count int, sum int, arr []int, num int) {
	if sum == 0 && count == 0 {
		cpy := make([]int, len(arr))
		copy(cpy, arr)
		ans = append(ans, cpy)
		return
	}
	if sum > 0 && count > 0 {
		for i := num; i <= 9; i++ {
			backtrace(count-1, sum-i, append(arr, i), i+1)
		}
	}

}
