package leetcode

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = 1
			continue
		}
		if nums[i] > nums[i-1] {
			result[i] = result[i-1]+1
		} else {
			result[i] = 1
		}
	}
	max:=0
	for i:=0;i<len(result);i++{
		if result[i]>max{
			max  = result[i]
		}
	}
	return max

}
