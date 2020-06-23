package leetcode

import "sort"

func nextPermutation(nums []int)  {
	// 从末尾两个开始交换 完成交换即输出 若一次都没完成直接返回升序
	for i:=len(nums)-1;i>=1;i--{
		if nums[i]>nums[i-1]{
			// 交换 循环遍历
			index:=i-1
			// 从 i到len-1挑选出一个最小的比nums[i-1]大的数字
			for j:=len(nums)-1;j>=i;j--{
				if nums[j]>nums[index]{
					index = j
					break
				}
			}
			// 交换
			if index!=i-1{

				nums[i-1],nums[index] = nums[index],nums[i-1]
				// 交换之后 i 到 len(num)-1需要升序排列
				slice:=nums[i:len(nums)]
				sort.Ints(slice)
				return
			}
		}
	}
	sort.Ints(nums)
	return
}