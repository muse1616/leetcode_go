package leetcode

func repeatedNTimes(A []int) int {
	hash := make(map[int]int, 10)
	for _, v := range A {
		hash[v]++
	}
	for k,v:=range hash{
		if v==len(A)/2{
			return k
		}
	}
	return 0
}
