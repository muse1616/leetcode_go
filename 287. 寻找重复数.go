package leetcode

func findDuplicate(nums []int) int {
	//	二分法 抽屉原理
	left, right := 1, len(nums)-1
	ans := -1
	for left <= right {
		middle := left + (right-left)/2
		//	统计数组中小于等于middle的数字
		// 因为 1~middle中如果没有重复数字 只有可能右middle个数字 如果大于middle个数字 说明重复数字在1~middle中
		count := 0
		for _, v := range nums {
			if v <= middle {
				count++
			}
		}
		if count <= middle {
			left = middle + 1
		} else if count > middle {
			right = middle - 1
			ans = middle
		}
	}
	return ans
}
