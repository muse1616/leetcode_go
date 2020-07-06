package leetcode

func productExceptSelf(nums []int) []int {
	//	左*右
	ans := make([]int, len(nums))
	//	先计算当前数字左边的乘积
	k := 1
	for i := 0; i < len(ans); i++ {
		ans[i] = k
		k *= nums[i]

	}
	//	再计算右边的 和 ans乘即可
	k = 1
	for i := len(ans) - 1; i >= 0; i-- {
		ans[i] *= k
		k *= nums[i]
	}
	return ans
}
