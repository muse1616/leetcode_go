package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	arr := make([]int, 0, m+n)
	i := 0
	j := 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			arr = append(arr, nums1[i])
			i++
		} else {
			arr = append(arr, nums2[j])
			j++
		}
	}
	if i == m {
		arr = append(arr, nums2[j:n]...)
	} else {
		arr = append(arr, nums1[i:m]...)
	}
	for i := 0; i < len(arr); i++ {
		nums1[i] = arr[i]
	}
}
