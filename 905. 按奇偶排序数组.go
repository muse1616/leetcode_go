package leetcode

func sortArrayByParity(A []int) []int {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	for _,v:=range A{
		if v%2==0{
			arr1 = append(arr1,v)
		}else {
			arr2 = append(arr2,v)
		}
	}

	return append(arr1, arr2...)
}
