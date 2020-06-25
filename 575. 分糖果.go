package leetcode

//排序后去重即可 暴力法
//func distributeCandies(candies []int) int {
//	sort.Ints(candies)
//	var tmp []int
//	tmp = append(tmp,candies[0])
//	for i:=1;i<len(candies);i++{
//		if candies[i]!=candies[i-1]{
//			tmp = append(tmp,candies[i])
//		}
//	}
//	if len(tmp)>=len(candies)/2{
//		return len(candies)/2
//	}else {
//		return len(tmp)
//	}
//}

//hash法
func distributeCandies(candies []int) int {
	hash := make(map[int]bool)
	types := 0
	for _, v := range candies {
		if types >= len(candies)/2 {
			return types
		}
		if _, ok := hash[v]; !ok {
			types++
			hash[v] = true
		}
	}
	return types
}
