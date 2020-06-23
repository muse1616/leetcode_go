package leetcode

import (
	"fmt"
	"sort"
)

// 类双指针解法
func fourSum(nums []int, target int) [][]int {
	// 结果四元组数组
	result := make([][]int, 0, 10)
	// 长度小于4 直接返回
	if len(nums) < 4 {
		return result
	}
	// 先升序排序
	sort.Ints(nums)
	// a+b+c+d=target 两头指针固定 中间两个指针夹逼
	// 指针 i j=i+1 k=l-1 l=len(nums)-1
	for i := 0; i <= len(nums)-4; i++ {

		// 去重
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}

		for l := len(nums) - 1; l >= i+3; l-- {
			// 去重
			if l < len(nums)-1 && nums[l] == nums[l+1] {
				continue
			}

			// 内部双指针两头扫描
			j := i + 1

			k := l - 1

			for {
				if j >= k {
					break
				}
				sum := nums[i] + nums[j] + nums[k] + nums[l]

				if sum > target {
					k--
					continue
				}

				if sum < target {
					j++
					continue
				}

				if sum == target {

					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					fmt.Println(i, j, k, l)
					for {
						if j < l && nums[j] == nums[j+1] {
							j++
						} else {
							break
						}
					}
					for {
						if k > i && nums[k] == nums[k-1] {
							k--
						} else {
							break
						}
					}
					j++
					k--
				}
			}

		}
	}
	return result
}
