package leetcode

func singleNumber(nums []int) []int {
	hash := make(map[int]int)
	for _, v := range nums {
		if hash[v] == 1 {
			delete(hash, v)
			continue
		}
		hash[v] = 1
	}
	var ans []int
	for k:=range hash{
		ans =append(ans,k)
	}
	return ans
}
