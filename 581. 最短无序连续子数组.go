package leetcode

import "sort"

func findUnsortedSubarray(nums []int) int {
	var tmp []int
	for i:=0;i<len(nums);i++{
		tmp = append(tmp,nums[i])
	}
	sort.Ints(tmp)
	start:=0
	last:=0
	for i:=0;i<len(tmp);i++{
		if nums[i] !=tmp[i]{
			start = i
			break
		}
	}
	for i:=len(tmp)-1;i>0;i--{
		if nums[i] !=tmp[i]{
			last = i
			break
		}
	}
	if last == start{
		return 0
	}
	return last-start+1
}
