package leetcode

import "strconv"

func summaryRanges(nums []int) []string {
	var result []string
	for i := 0; i < len(nums); {
		j := i
		for ; j < len(nums)-1; j++ {
			if nums[j+1] == nums[j]+1 {
				continue
			} else {
				break
			}
		}
		//孤立数字
		if i == j {
			result = append(result, strconv.Itoa(nums[i]))
		} else {
			//	此区间为 i -- j
			str := strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j])
			result = append(result, str)
		}
		i = j + 1

	}

	return result
}
