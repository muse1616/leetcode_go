package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums2)==0{
		return []int{}
	}
	hash := make(map[int]int, len(nums2))
	for i := 0; i < len(nums2)-1; i++ {
		f:=false
		for j := i + 1;j<len(nums2);j++{
			if nums2[j]>nums2[i]{
				hash[nums2[i]] = nums2[j]
				f = true
				break
			}
		}
		if f==false{
			hash[nums2[i]]=-1
		}

	}
	hash[nums2[len(nums2)-1]] = -1

	var res	[]int
	for i:=0;i<len(nums1);i++{
		res = append(res,hash[nums1[i]])
	}
	return res
}
