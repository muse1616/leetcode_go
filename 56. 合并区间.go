package leetcode

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	//先排序
	if len(intervals) <= 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//两两合并
	for i := 0; i < len(intervals); {
		j := i + 1
		for ; j < len(intervals); {
			if intervals[i][1] < intervals[j][0] || intervals[i][0] > intervals[j][1] {
				break
			} else {
				intervals[i][0] = int(math.Min(float64(intervals[i][0]), float64(intervals[j][0])))
				intervals[i][1] = int(math.Max(float64(intervals[i][1]), float64(intervals[j][1])))
				intervals = append(intervals[:j], intervals[j+1:]...)
			}
		}
		i = j
	}
	return intervals
}
