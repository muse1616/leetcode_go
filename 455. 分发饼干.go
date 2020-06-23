package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	count := 0
	//	先对两数组排序
	sort.Ints(g)
	sort.Ints(s)
	for i, j := 0, 0; i < len(g) && j < len(s); i++ {
		if g[i] > s[len(s)-1] {
			//	当前需要的饼干比最大的饼干还大直接返回
			return count
		}
		//贪心找到符合当前胃口的最小饼干
		for g[i] > s[j] {
			j++
		}
		j++
		count++
	}
	return count
}
