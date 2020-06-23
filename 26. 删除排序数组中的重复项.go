package leetcode

func removeDuplicates(nums []int) int {
	count:=1
	for _,val:=range nums{
		if val!= nums[count-1]{
			nums[count] = val
			count++
		}
	}

	return count
}
