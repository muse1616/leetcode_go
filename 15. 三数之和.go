package leetcode
func threeSum(nums []int) [][]int {
	//
	if nums == nil || len(nums) < 3 {
		return nil
	}
	// 排序
	sort.Ints(nums)
	// 结果三元组
	result := make([][]int, 0, 5)

	for i:=0;i<len(nums)-2;i++{
		if nums[i]>0{
			break
		}
		// 去重
		if i>=1&&nums[i-1]==nums[i]{
			continue
		}
		j:=i+1
		k:=len(nums)-1

		for{
			if j>=k{
				break
			}
			sum:=nums[i]+nums[j]+nums[k]
			if sum>0{
				k--
				continue
			}
			if sum<0{
				j++
				continue
			}
			if sum==0{
				result = append(result, []int{nums[i],nums[j],nums[k]})
				for{
					if k>i&&nums[k-1] ==nums[k]{
						k--
					}else {
						break
					}
				}
				for{
					if j<i&&nums[j]==nums[j+1]{
						j++
					}else {
						break
					}
				}
				j++
				k--
			}

		}

	}

	return result
}