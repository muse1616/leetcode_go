package leetcode

import "math"

//单调栈实现
func maximalRectangle(matrix [][]byte) int {

	if matrix == nil || len(matrix) == 0 {
		return 0
	}
	//保存最终结果
	max := 0
	//行数，列数
	m, n := len(matrix), len(matrix[0])
	//高度数组（统计每一行中1的高度）
	height := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//每一行去找1的高度
			//如果不是‘1’，则将当前高度置为0
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				//是‘1’，则将高度加1
				height[j] = height[j] + 1
			}
		}
		//更新最大矩形的面积
		max = int(math.Max(float64(max), float64(maxRectangle(height))))
	}
	return max
}

//使用84题的结果
func maxRectangle(heights []int) int {
	// 高度首尾添加负数 形成单调栈判断
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)

	size := len(heights)

	// 递增栈
	s := make([]int, 1, size)

	result := 0

	for i := 1; i < len(heights); {
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
