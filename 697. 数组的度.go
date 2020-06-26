package leetcode

import (
	"math"
)

func findShortestSubArray(nums []int) int {
	//	找到度的元素
	hash := make(map[int]int)
	for _, v := range nums {
		if _, ok := hash[v]; ok {
			hash[v]++
		} else {
			hash[v] = 1
		}
	}
	//找到度
	max := 0
	for _, v := range hash {
		if v > max {
			max = v
		}
	}
	var result []int
	for k, v := range hash {
		if v == max {
			result = append(result, k)
		}
	}
	//子串长度是result数组中第一次出现和最后一次出现的最小值
	minRange := 50001
	for _,v:=range result{
		//	找第一次出现和最后一次出现
		var tmp int
		i:=0
		j:=len(nums)-1
		for ;i<len(nums);i++{
			if nums[i] == v{
				break
			}
		}
		for ;j>0;j--{
			if nums[j] == v{
				break
			}
		}
		tmp = j-i+1
		minRange =int(math.Min(float64(minRange),float64(tmp)))
	}

	return minRange
}
