package leetcode

var result [][]int

func combine(n int, k int) [][]int {
	result = make([][]int, 0)
	// idx为剩余的数字
	backtrace([]int{}, n, n, k)
	return result

}

// 每次遇到一个数字
// 1. len(arr)+idx > k 放或不放
// 2. len(arr)+idx  = k 必须放
// 3. len(arr)+idx  < k 不成立返回
func backtrace(arr []int, idx int, n int, k int) {
	if idx == 0 {
		if len(arr) == k {
			cpy := make([]int, len(arr))
			copy(cpy, arr)
			result = append(result, cpy)
		}
		return
	}
	if len(arr)+idx >= k {
		//不放
		backtrace(arr, idx-1, n, k)
		//放
		arr = append(arr, n-idx+1)
		backtrace(arr, idx-1, n, k)
	}
	if len(arr)+idx < k {
		return
	}
}
