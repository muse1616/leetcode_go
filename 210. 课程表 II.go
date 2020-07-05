package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 拓扑排序
	//1. 入度表
	indegrees := make([]int, numCourses)
	//2. 邻接表 i->j
	adjaceny := make([][]int, numCourses)
	//3. 临时队列
	queue := make([]int, 0)
	// 入度表邻接表预处理
	for i := 0; i < len(prerequisites); i++ {
		// 入度+1
		indegrees[prerequisites[i][0]]++
		// 邻接表
		adjaceny[prerequisites[i][1]] = append(adjaceny[prerequisites[i][1]], prerequisites[i][0])

	}
	// 将入度为0的元素全部压入队列
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	var res []int
	for len(queue) > 0 {
		//	出列
		tmpNode := queue[0]
		queue = queue[1:]
		res = append(res,tmpNode)
		numCourses--
		//	在邻接表中删除需要该元素作为前置课程的元素 即 j的入度-1
		for i := 0; i < len(adjaceny[tmpNode]); i++ {
			indegrees[adjaceny[tmpNode][i]]--
			if indegrees[adjaceny[tmpNode][i]] == 0 {
				queue = append(queue, adjaceny[tmpNode][i])
			}
		}

	}
	if numCourses!=0{
		return nil
	}
	return res

}
