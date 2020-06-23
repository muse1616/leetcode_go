package leetcode

import "math"

func threeSumClosest(nums []int, target int) int {
	if nums==nil||len(nums)<3{
		return 0
	}
	min:=int(math.Abs(float64(nums[0]+nums[1]+nums[2]-target)))
	result:=nums[0]+nums[1]+nums[2]
	for i:=0;i<len(nums);i++{
		for j:=i;j<len(nums);j++{
			for k:=j;k<len(nums);k++{
				if i!=k&&i!=j&&k!=j{
					if math.Abs(float64(nums[i]+nums[j]+nums[k]-target))<float64(min){
						min = int(math.Abs(float64(nums[i]+nums[j]+nums[k]-target)))
						result = nums[i]+nums[j]+nums[k]
					}
				}
			}
		}
	}
	return result
}
