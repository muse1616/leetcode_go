package leetcode

func findPeakElement(nums []int) int {
	//规律一：如果nums[i] > nums[i+1]，则在i之前一定存在峰值元素
	//规律二：如果nums[i] < nums[i+1]，则在i+1之后一定存在峰值元素
	// 总结 大的那一半肯定存在峰值
	// 解释: 如果 nums[i+1]>nums[i]
	// 那么已知nums[i+1]>nums[i] 第一种情况是nums[i+2]<nums[i+1] 那么nums[i+1]就是峰值
	// 第二种情况是 nums[i+2]>nums[i+1] 如果后面一直递增 则最后一个元素是峰值 因为nums[n]~负无穷
	// 同理nums[i]>nums[i+1]为对称情况

	i, j := 0, len(nums)-1
	for i < j {
		mid := i + (j-i)/2
		if nums[mid] > nums[mid+1] {
			// 这里不能加1 因为没有排除mid不是峰值的情况
			j = mid
		} else {
			// 这里需要加1 因为已经知道了mid不是峰值了 不然无限循环
			i = mid + 1
		}
	}
	return i
}
