package leetcode

//func rotate(nums []int, k int) {
//	copy(nums, append(nums[len(nums)-k%len(nums):], nums[0:len(nums)-k%len(nums)]...))
//}
func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
