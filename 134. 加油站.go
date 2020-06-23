package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	start := -1
	//遍历每个点 以每个点作为起点
	for i := 0; i < len(gas); i++ {
		//初始油量
		sum := gas[i]
		j := i
		for {
			if sum-cost[j] < 0 {
				break
			}
			//j为最后加油站 直接赋值为0
			if j == len(gas)-1 {
				sum = sum + gas[0] - cost[j]
				if sum >= 0 && i == 0 {
					return 0
				}
			} else {
				sum = sum + gas[j+1] - cost[j]
			}

			if sum < 0 {
				break
			}
			// 返回终点
			if j+1 == i {
				return i
			}
			if j == len(gas)-1 {
				j = 0
			} else {
				j++
			}
		}
	}

	return start
}