package leetcode

// 0-红色 1-白色 2-蓝色
func sortColors(nums []int) {
	//hash := make(map[int]int, 3)
	//for _, v := range nums {
	//	hash[v]++
	//}
	//result := make([]int, 0, len(nums))
	//for i := 0; i < hash[0]; i++ {
	//	result = append(result, 0)
	//}
	//for i := 0; i < hash[1]; i++ {
	//	result = append(result, 1)
	//}
	//for i := 0; i < hash[2]; i++ {
	//	result = append(result, 2)
	//}
	//copy(nums, result)
	left := 0
	right := len(nums) - 1
	for i := 0; i <= right; i++ {
		// 找到0 交换到最左边 即左指针
		if nums[i] == 0 {
			// 交换位置
			nums[i], nums[left] = nums[left], nums[i]
			// 此时i不需要-- 因为换回来的0也无所谓 换回来1后面会处理
			left++
		}
		// 找到2 交换到右边
		if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			// 此时i需要-1 因为换到了右边 有可能换回来的依然是2
			i--
			right--
		}
	}

}
