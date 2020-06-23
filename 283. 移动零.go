package leetcode

func moveZeroes(nums []int)  {
	l:=len(nums)
	for i:=0;i<len(nums);{
		if nums[i] == 0{
			nums = append(nums[:i],nums[i+1:]...)
			continue
		}
		i++
	}
	arr:=make([]int,l-len(nums),l-len(nums))
	nums = append(nums,arr...)
}