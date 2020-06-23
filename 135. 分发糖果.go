package leetcode


func candy(ratings []int) int {
	arr := make([]int, len(ratings))
	arr[0] = 1
	for i := 1; i < len(ratings); i++ {
		//比前一个孩子分数高
		if ratings[i] > ratings[i-1] {
			//多发一颗糖
			arr[i] = arr[i-1] + 1
		} else if ratings[i]==ratings[i-1]{
			if arr[i-1]>1{
				arr[i] = 1
			}else {
				arr[i] = arr[i-1]
			}
		} else {
			//	比之前的孩子分数少
			if arr[i-1] > 1 {
				arr[i] = 1
			} else if arr[i-1] == 1 {
				arr[i] = arr[i-1]
				arr[i-1]++
				for j := i - 1; j >= 1; j-- {
					if ratings[j]==ratings[j-1]{
						break
					}
					if ratings[j] < ratings[j-1] && arr[j] >= arr[j-1] {
						arr[j-1]++
					}else {
						break
					}
				}
			}

		}
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
