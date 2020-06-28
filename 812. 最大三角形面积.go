package leetcode

import "math"

//S=(1/2)abs[(x1y2+x2y3+x3y1-x1y3-x2y1-x3y2)] 面积公式
func largestTriangleArea(points [][]int) float64 {
	var maxArea float64
	maxArea = 0
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {

				nowArea := 0.5 * math.Abs(float64(points[i][0]*points[j][1]+points[j][0]*points[k][1]+points[k][0]*points[i][1]-points[i][0]*points[k][1]-points[j][0]*points[i][1]-points[k][0]*points[j][1]))
				maxArea = math.Max(maxArea, nowArea)
			}
		}
	}

	return maxArea
}
