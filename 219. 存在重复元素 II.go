package leetcode

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			if int(math.Abs(float64(m[nums[i]])-float64(i))) <= k {
				return true
			}else {
				m[nums[i]] = i
			}
		} else {
			m[nums[i]] = i
		}
	}

	return false
}
