package leetcode

import "math"

func constructRectangle(area int) []int {
	//gap := int(math.Abs(float64(area - 1)))
	//rec := make([]int, 2)
	//rec[0] = area
	//rec[1] = 1
	//for i := 1; i <= area/2+1; i++ {
	//	if i*(area/i) == area && int(math.Abs(float64(i-area/i))) <= gap && i >= area/i {
	//		gap = int(math.Abs(float64(i - area/i)))
	//		rec[0] = i
	//		rec[1] = area / i
	//
	//	}
	//}
	//return rec
	sqr := int(math.Sqrt(float64(area)))
	for area%sqr != 0 {
		sqr--
	}
	return []int{area / sqr, sqr}
}
