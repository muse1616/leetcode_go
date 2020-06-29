package leetcode

import (
	"math"
)

func robotSim(commands []int, obstacles [][]int) int {
	type point struct {
		x int
		y int
	}
	maxDistance := 0
	//	方向向量
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	//	障碍哈希 hash[x] = y
	hash := make(map[point]bool, 10)
	for i := 0; i < len(obstacles); i++ {
		tmp := point{obstacles[i][0], obstacles[i][1]}
		hash[tmp] = true
	}
	// 起始点
	startPointX := 0
	startPointY := 0
	// 表示方向坐标的index
	startDirection := 0

	for i := 0; i < len(commands); i++ {
		//向当前方向移动指令
		if commands[i] >= 1 && commands[i] <= 9 {
			//循环移动 判断前方是否是障碍
			for j := 1; j <= commands[i]; j++ {
				//遇到障碍物
				p := point{startPointX + dx[startDirection]*1, startPointY + dy[startDirection]*1}
				if hash[p] == true && len(hash) != 0 {
					break
				} else {
					startPointX += dx[startDirection] * 1
					startPointY += dy[startDirection] * 1
					maxDistance = int(math.Max(float64(maxDistance), float64(startPointX*startPointX+startPointY*startPointY)))
				}
			}
		} else if commands[i] == -2 {
			//左转90度
			if startDirection == 0 {
				startDirection = 3
			} else if startDirection == 1 {
				startDirection = 0
			} else if startDirection == 2 {
				startDirection = 1
			} else if startDirection == 3 {
				startDirection = 2
			}
		} else if commands[i] == -1 {
			//右转90度
			if startDirection == 0 {
				startDirection = 1
			} else if startDirection == 1 {
				startDirection = 2
			} else if startDirection == 2 {
				startDirection = 3
			} else if startDirection == 3 {
				startDirection = 0
			}
		}
	}

	return maxDistance
}
