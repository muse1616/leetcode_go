package leetcode

func findErrorNums(nums []int) []int {
	hash := make(map[int]int)
	//重复的元素
	repeat := 0
	for _, v := range nums {
		if _, ok := hash[v]; !ok {
			hash[v] = 1
		} else {
			hash[v]++
			repeat = v
		}
	}
	for i:=1;i<=len(nums);i++{
		if _,ok:=hash[i];!ok{
			return []int{repeat,i}
		}
	}
	return nil

}
