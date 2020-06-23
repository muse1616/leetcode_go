package leetcode

import "math"

func largestRectangleArea(heights []int) int {

	// 高度首尾添加负数 形成单调栈判断
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)

	size := len(heights)

	// 递增栈
	s:=make([]int,1,size)

	result:=0

	for i:=1;i < len(heights); {
		// 递增则入栈
		if heights[s[len(s)-1]] < heights[i] {
			s = append(s, i)
			i++
			continue
		}
		// 面积计算 此时栈中存放的序号已经保证了递增 如果当前比较的序号比栈顶小
		// 说明这个数字到左边界的距离里都是比左边界高的 面积即左边界乘以到左边界的距离
		result = int(math.Max(float64(result), float64(heights[s[len(s)-1]]*(i-s[len(s)-2]-1))))
		s = s[:len(s)-1]
	}
	return result


}


