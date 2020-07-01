package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	//连续出现次数
	count := 1
	for i := 1; i < len(nums); {
		// 等于前一个数字
		if nums[i] == nums[i-1] {
			//出现次数加1
			count++
			//出现三次该数字
			if count == 3 {
				//	删除改该元素
				nums = append(nums[:i], nums[i+1:]...)
				//出现次数变为2
				count = 2
			} else {
				i++
			}
		} else {
			//重置出现次数
			count = 1
			i++
		}
	}

	return len(nums)
}
