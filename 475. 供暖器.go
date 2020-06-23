package leetcode

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	//先排序 坑
	sort.Ints(houses)
	sort.Ints(heaters)
	//最大距离
	maxDistance := 0
	//遍历房间
	for _, house := range houses {
		left := 0
		right := len(heaters) - 1
		houseRes := 0
		// 二分查找
		for left < right {
			middle := left + (right-left)/2
			if heaters[middle] < house {
				left = middle + 1
			} else {
				right = middle
			}
		}

		if heaters[left] == house {
			houseRes = 0
		} else if heaters[left] < house { // 供暖器在左侧
			houseRes = house - heaters[left]
		} else { // 供暖器在右侧
			if left == 0 {
				houseRes = heaters[left] - house
			} else {
				houseRes = int(math.Min(float64(heaters[left]-house), float64(house-heaters[left-1])))
			}
		}

		maxDistance = int(math.Max(float64(maxDistance), float64(houseRes)))
	}

	return maxDistance
}
