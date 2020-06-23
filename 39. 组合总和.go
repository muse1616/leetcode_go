package leetcode

import (
	"reflect"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	// 结果集
	result := make([][]int, 0, 10)
	backtracking(candidates, []int{}, &result, 0, target)
	return result
}
func backtracking(candidates []int, item []int, result *[][]int, count int, target int) {
	if count >= len(candidates) {
		return
	}

	// target减为0直接返回结果
	if target == 0 {
		// result里面没有才加
		b := false

		for i := 0; i < len(*result); i++ {
			if reflect.DeepEqual((*result)[i], item) {
				b = true
			}
		}
		if b == false {
			*result = append(*result, item)
		}

		return
	}
	// 结果小于0 也结束回溯
	if target < 0 {
		return
	}
	// 两种情况 一种是继续加自己 一种是加下一个
	// 1.加下一个数
	newItem := make([]int, len(item))
	if count < len(candidates)-1 {

		copy(newItem, item)
		newItem = append(newItem, candidates[count+1])
		backtracking(candidates, newItem, result, count+1, target-candidates[count+1])

	}

	// 2.加自己
	newItem = make([]int, len(item))
	copy(newItem, item)
	newItem = append(newItem, candidates[count])
	backtracking(candidates, newItem, result, count, target-candidates[count])

	// 3.不加
	newItem = make([]int, len(item))
	copy(newItem, item)
	backtracking(candidates, newItem, result, count+1, target)

}
