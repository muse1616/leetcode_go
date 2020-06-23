package leetcode

import "math/rand"

//func singleNumber(nums []int) int {
//	m := make(map[int]int, len(nums))
//	for _, v := range nums {
//		m[v]++
//	}
//
//	for k, v := range m {
//		if v == 1 {
//			return k
//		}
//	}
//
//}
// 异或算法
func singleNumber(nums []int) int {
	target:=0
	for _,v:=range nums{
		target ^= v
	}
	return target

}