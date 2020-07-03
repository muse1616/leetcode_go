package leetcode

func singleNumber(nums []int) int {
	//hash:=make(map[int]int,0)
	//for _,v:=range nums{
	//	hash[v]++
	//}
	//for k,v:=range hash{
	//	if v==1{
	//		return k
	//	}
	//}
	//return 0
	a, b := 0, 0
	for _, num := range nums {
		a = a ^ num&^b // num先对b按位清除；再与a异或
		b = b ^ num&^a // num先对b按位清除；再与a异或
	}
	return a

}
