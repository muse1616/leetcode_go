package leetcode

import "math"

// s1+s2-重叠
// 正方形1：左下(A,B) 右上(C,D)
// 正方形2：左下(E,F) 右上(G,H)
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	s1 := (C - A) * (D - B)
	s2 := (G - E) * (H - F)
	//	1.未发生重叠 上下左右四种情况
	if H <= B || E >= C || G <= A || F >= D {
		return s1 + s2
	}
	x1 := math.Max(float64(A), float64(E))
	y1 := math.Max(float64(B), float64(F))
	x2 := math.Min(float64(C), float64(G))
	y2 := math.Min(float64(D), float64(H))
	return s1 + s2 - int((x2-x1)*(y2-y1))
}
