package leetcode

import "math"

func minSubArrayLen(s int, nums []int) int {
	//if len(nums) == 0{
	//	return 0
	//}
	//minLen:=math.MaxInt64
	//for i:=0;i<len(nums);i++{
	//	sum:=nums[i]
	//	if sum>=s{
	//		return 1
	//	}
	//	for j:=i+1;j<len(nums);j++{
	//		sum+=nums[j]
	//		if sum>=s{
	//			minLen = int(math.Min(float64(minLen),float64(j-i+1)))
	//			break
	//		}
	//	}
	//}
	//if minLen == math.MaxInt64{
	//	return 0
	//}
	//return minLen

	if len(nums) == 0 {
		return 0
	}
	ans := math.MaxInt32
	// 双指针 滑动窗口
	start, end := 0, 0
	sum := 0
	for end < len(nums) {
		// 不停地将end右移 直到大于所给数
		sum += nums[end]
		for sum >= s {
			// 此时start开始右移 查看能不能减去一个数字
			ans = int(math.Min(float64(ans), float64(end-start+1)))
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans

}
