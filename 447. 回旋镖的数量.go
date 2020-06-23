package leetcode

func numberOfBoomerangs(points [][]int) int {
	count := 0

	for i1, v1 := range points {
		hash := make(map[int]int)
		for i2, v2 := range points {
			if i1 != i2 {
				distance := helper(v1, v2)
				if val, ok := hash[distance]; ok {
					count += val * 2
				}
				hash[distance]++
			}
		}
	}

	return count
}

// 返回两点距离
func helper(i []int, j []int) int {
	return (i[0]-j[0])*(i[0]-j[0]) + (i[1]-j[1])*(i[1]-j[1])
}
