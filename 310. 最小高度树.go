package leetcode

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	//	bfs 度图和图
	graph := make([][]int, n)
	degrees := make([]int, n)
	for _, edge := range edges {
		//两个点的度都加1
		degrees[edge[0]]++
		degrees[edge[1]]++
		//	初始化图
		if graph[edge[0]] == nil {
			graph[edge[0]] = []int{}
		}
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		//	注意双向连接
		if graph[edge[1]] == nil {
			graph[edge[1]] = []int{}
		}
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	//	将度为1的节点如队列
	queue := make([]int, 0)
	for k, v := range degrees {
		if v == 1 {
			queue = append(queue, k)
		}
	}
	var ans []int
	for len(queue) != 0 {
		l := len(queue)
		ans = []int{}
		for i := 0; i < l; i++ {
			//度减1
			degrees[queue[i]] --
			//	相邻节点-1
			for _, neighbour := range graph[queue[i]] {
				degrees[neighbour] --
				if degrees[neighbour] == 1 {
					queue = append(queue, neighbour)
				}
			}
			// 此时如果ans不被上方的重置语句所重置 说明队列中的元素已经是无法去掉的了
			ans = append(ans, queue[i])
		}
		//	出队
		queue = queue[l:]
	}

	return ans
}
