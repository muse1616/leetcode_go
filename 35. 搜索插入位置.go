package leetcode
func searchInsert(nums []int, target int) int {
	// 二分查找插入位置
	begin,end:=0,len(nums)-1
	mid:=(begin+end)/2
	for begin<=end {
		mid=(begin+end)/2
		if nums[mid]>target{
			end = mid-1
		}else if nums[mid]<target{
			begin = mid+1
		}else{
			return mid
		}
	}
	// 未找到的情况 确定插入位置
	if nums[mid]<target{
		return mid+1
	}
	return mid
}