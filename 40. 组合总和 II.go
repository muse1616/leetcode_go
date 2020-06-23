package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	// 排序
	sort.Ints(candidates)
	// 答案
	var ans [][]int

	backtrack(candidates, target, 0, []int{}, &ans)
	return ans
}

func backtrack(candidates []int, target, index int, path []int, ans *[][]int) {
	// 使用target减当前数的模式 若target==0就结束回溯
	if target == 0 {
		*ans = append(*ans, path)
		return
	}

	// 跳过重复项
	for i := index; i < len(candidates); i++ {
		rawI := i
		for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
			i++
		}
		// 剪枝
		if candidates[i] > target {
			return
		}

		newPath := make([]int, len(path))
		copy(newPath, path)
		newPath = append(newPath, candidates[i])
		backtrack(candidates, target-candidates[i], rawI+1, newPath, ans)

	}

}
