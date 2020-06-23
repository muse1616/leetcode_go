package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	var result []int
	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] ++
	}
	for i := 0; i < len(nums2); i++ {
		if m[nums2[i]] != 0 {
			result = append(result, nums2[i])
			m[nums2[i]]--
		}
	}
	return result
}