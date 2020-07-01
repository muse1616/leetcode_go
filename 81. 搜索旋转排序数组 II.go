package leetcode

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	// 找到旋转点
	i := 0
	for ; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			break
		}
	}
	begin := 0
	end := 0
	if target < nums[0] {
		// 在后半部分二分
		begin = i + 1
		end = len(nums) - 1

	} else {
		// 前半部分二分
		begin = 0
		end = i
	}

	for {
		if begin > end {
			break
		}
		mid := (begin + end) / 2
		if target > nums[mid] {
			begin = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else if target == nums[mid] {
			return true
		}
	}

	return false
}
