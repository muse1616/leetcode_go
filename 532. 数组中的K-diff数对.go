package leetcode

import "sort"

func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	sum:=0
	//双指针
	for i:=0;i<len(nums)-1;i++{
		if i>=1&&nums[i] == nums[i-1] {
			continue
		}
		for j:=i+1;j<len(nums);j++{
			if nums[j] == nums[j-1]&& j>i+1{
				continue
			}
			if nums[j]-nums[i]>k{
				break
			}
			if nums[j]-nums[i]==k{
				sum++
			}
		}
	}
	return sum
}
