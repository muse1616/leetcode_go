package leetcode

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	//	初始化dp
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			//	终点 终点可能为正数 这样到终点时 1 点生命值即可 也可能为负数 这样需要 1-这个负数的生命值 取大的
			if i == m-1 && j == n-1 {
				dp[i][j] = int(math.Max(float64(1), float64(1-dungeon[i][j])))
			} else if j == n-1 {
				//	最后一列 同终点 1点生命值或是后一步需要的生命值减去这次扣得生命值
				dp[i][j] = int(math.Max(float64(1), float64(dp[i+1][j]-dungeon[i][j])))
			} else if i == m-1 {
				//	最后一行
				dp[i][j] = int(math.Max(float64(1), float64(dp[i][j+1]-dungeon[i][j])))
			} else {
				//中间部分 需要预留出下一步需要的血量 有两种走法 下或是右 即选取下方或右方最小的值减去当前格子的血
				dp[i][j] = int(math.Max(1, math.Min(float64(dp[i+1][j]), float64(dp[i][j+1]))-float64(dungeon[i][j])))
			}
		}
	}

	//左上角
	return dp[0][0]
}
