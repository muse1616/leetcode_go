package leetcode

func findDisappearedNumbers(nums []int) []int {
	var result []int
	hash := make(map[int]bool, len(result))
	for i := 1; i <= len(nums); i++ {
		hash[i] = false
	}
	for i:=0;i<len(nums);i++{
		hash[nums[i]] = true
	}
	for k, v := range hash{
		if v == false{
			result = append(result,k)
		}
	}

	return result
}
