package leetcode

func majorityElement(nums []int) []int {
	var ans []int
	if len(nums) == 0 {
		return nil
	}
	// 大于三分之一的人最多只有两个
	candidate1 := nums[0]
	count1 := 0
	candidate2 := nums[0]
	count2 := 0
	// 摩尔投票法
	// 投票阶段
	for _, v := range nums {
		if candidate1 == v {
			count1++
			continue
		}
		if candidate2 == v {
			count2++
			continue
		}
		//	更换候选人
		if count1 == 0 {
			candidate1 = v
			count1++
			continue
		}
		if count2 == 0 {
			candidate2 = v
			count2++
			continue
		}
		//	两个都不同 抵消票数
		count2--
		count1--
	}
	// 计数阶段 此时count>0的有可能是大于1/3 验证
	count1 = 0
	count2 = 0
	for _, v := range nums {
		if candidate1 == v {
			count1++
		} else if candidate2 == v {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		ans = append(ans, candidate1)
	}
	if count2 > len(nums)/3 {
		ans = append(ans, candidate2)
	}

	return ans
}
