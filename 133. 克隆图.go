package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	// 空图
	if node == nil {
		return nil
	}
	// hash[原图节点的val] = 连接点
	hash := make(map[int]*Node)
	// 模拟bfs遍历队列
	buf := []*Node{node}
	//结果
	newNode := Node{Val: node.Val, Neighbors: []*Node{}}
	hash[node.Val] = &newNode

	for len(buf) > 0 {
		//临时队列
		var t []*Node
		//遍历队列
		for i := 0; i < len(buf); i++ {
			// 取当前节点
			n := buf[i]
			//遍历改节点的邻接点
			for j := 0; j < len(n.Neighbors); j++ {
				//该点是否遍历过
				if v, ok := hash[n.Neighbors[j].Val];!ok{
					newN := Node{Val: n.Neighbors[j].Val, Neighbors: []*Node{}}
					hash[n.Neighbors[j].Val] = &newN
					t = append(t, n.Neighbors[j])
					hash[n.Val].Neighbors = append(hash[n.Val].Neighbors, &newN)
				}else {
					//遍历过
					hash[n.Val].Neighbors = append(hash[n.Val].Neighbors, v)
				}
			}
		}
		// 更新队列
		buf = t
	}
	return &newNode

}
