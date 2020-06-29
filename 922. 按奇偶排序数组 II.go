package leetcode

func sortArrayByParityII(A []int) []int {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	for _, v := range A {
		if v%2 == 0 {
			arr2 = append(arr2, v)
		} else {
			arr1 = append(arr1, v)
		}
	}
	result := make([]int, 0)
	for i := 0; i < len(A); i++ {
		if i%2 == 0 {
			result = append(result, arr2[i/2])
		} else {
			result = append(result, arr1[i/2])
		}
	}
	return result
}
