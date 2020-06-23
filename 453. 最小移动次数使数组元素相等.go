package leetcode

func minMoves(nums []int) int {
//	长度为n 当前数组和为sum 需要移动m次 移动到每个数字是x 及 sum+m*（n-1） = x*n
//	设最小值为minVal 即minVal = x-m x=m+minVal
//	m = sum-minVal * n
	return sum(nums) - min(nums)*len(nums)
}
func sum(nums []int)int{
	sumVal :=0
	for i:=0;i<len(nums);i++{
		sumVal += nums[i]
	}
	return sumVal
}
func min(nums []int)int  {
	minVal:=nums[0]
	for i:=1;i<len(nums);i++{
		if nums[i] < minVal{
			minVal = nums[i]
		}
	}
	return minVal
}