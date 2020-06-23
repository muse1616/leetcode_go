package leetcode

//func twoSum(numbers []int, target int) []int {
//	for i := 0; i < len(numbers); i++ {
//		for j := i + 1; j < len(numbers); j++ {
//			if numbers[i]+numbers[j] == target {
//				return []int{i+1, j+1}
//			}
//		}
//	}
//	return []int{}
//}

func twoSum(numbers []int, target int) []int {
	left:=0
	right:=len(numbers)-1
	for left<right{
		if numbers[left]+numbers[right]==target{
			return []int{left+1,right+1}
		}else if numbers[left]+numbers[right]<target{
			left++
		}else {
			right--
		}
	}
	return []int{}
}