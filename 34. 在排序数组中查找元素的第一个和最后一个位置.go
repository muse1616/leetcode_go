package leetcode
func searchRange(nums []int, target int) []int {
	// 升序排列
	begin:=0
	end:=len(nums)-1
	for {
		if begin>end{
			break
		}
		mid:=(begin+end)/2
		if target<nums[mid]{
			end = mid -1
		}else if target>nums[mid] {
			begin = mid+1
		}else if target==nums[mid]{
			// 前后搜索
			i:=mid
			for ;i>=begin;i--{
				if nums[i]!=target{
					break
				}
			}
			j:=mid
			for ;j<=end;j++{
				if nums[j]!=target{
					break
				}
			}
			return []int{i+1,j-1}

		}

	}
	return []int{-1,-1}

}