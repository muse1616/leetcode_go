package leetcode

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	//不重叠 左 右 上 下
	return !(rec1[2]<=rec2[0]||rec1[0]>=rec2[2]||rec1[1]>=rec2[3]||rec1[3]<=rec2[1])
}
