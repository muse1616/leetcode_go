package leetcode

// 线段树
func getSkyline(buildings [][]int) [][]int {
	len := len(buildings)
	if len == 0 {
		return nil
	}
	return segment(buildings, 0, len-1)
}

func segment(buildings [][]int, l int, r int) [][]int {
	// 创建返回值
	var res [][]int
	// 结束条件
	if l == r {
		return [][]int{{buildings[l][0], buildings[l][2]}, {buildings[l][1], 0}}
	}
	// 取中间值
	mid := l + (r-l)/2
	// 左递归
	left := segment(buildings, l, mid)
	// 右递归
	right := segment(buildings, mid+1, r)
	// 左右合并
	// 创建left 和 right 的索引位置
	var m int = 0
	var n int = 0
	// 创建 left 和 right 的目前高度
	var lpre int = 0
	var rpre int = 0
	for m < len(left) || n < len(right) {
		// 一边遍历完，则全部添加另一边
		if m >= len(left) {
			res = append(res, right[n])
			n++
		} else if n >= len(right) {
			res = append(res, left[m])
			m++
		} else { // swip line
			if left[m][0] < right[n][0] {
				if left[m][1] > rpre {
					res = append(res, left[m])
				} else if lpre > rpre {
					res = append(res, []int{left[m][0], rpre})
				}
				lpre = left[m][1]
				m++
			} else if right[n][0] < left[m][0] {
				if right[n][1] > lpre {
					res = append(res, right[n])
				} else if rpre > lpre {
					res = append(res, []int{right[n][0], lpre})
				}
				rpre = right[n][1]
				n++
			} else { // left 和 right横坐标相等
				if left[m][1] >= right[n][1] && left[m][1] != max(lpre, rpre) {
					res = append(res, left[m])
				} else if left[m][1] <= right[n][1] && right[n][1] != max(lpre, rpre) {
					res = append(res, right[n])
				}
				lpre = left[m][1]
				rpre = right[n][1]
				m++
				n++
			}
		}
	}
	return res
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}
