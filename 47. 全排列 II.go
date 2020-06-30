package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	var result [][]int
	//是为了吧相同的元素放在一起 进行剪枝
	sort.Ints(nums)
	backtrace(nums, 0, &result)
	return result
}
func backtrace(nums []int, i int, result *[][]int) {
	n := len(nums)
	//完成回溯
	if i == len(nums)-1 {
		n := make([]int, len(nums))
		copy(n, nums)
		*result = append(*result, n)
	}
	// nums[0:i]是已经决定的部分，nums[i:]是待决定部分，同时待选元素也都在nums[i:]
	for k := i; k < n; k++ {
		// 跳过重复数字 已经排序了
		if k != i && nums[k] == nums[i] {
			continue
		}
		// 交换
		nums[k], nums[i] = nums[i], nums[k]
		// 回溯
		backtrace(nums, i+1, result)
	}
	// 还原状态
	for k := n - 1; k > i; k-- {
		nums[i], nums[k] = nums[k], nums[i]
	}
}
