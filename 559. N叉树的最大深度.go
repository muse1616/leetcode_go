package leetcode

import "math"

type Node struct {
	Val      int
	Children []*Node
}

//只要找出当前节点的子节点中最大的深度 加1就是最大深度
func maxDepth(root *Node) int {
	//叶子 0
	if root == nil {
		return 0
	}
	//默认值 0 即子节点为叶子节点的节点深度 0+1
	var depth int
	//遍历子节点
	for _, v := range root.Children {
		//深度是当前最大深度和叶子最大深度中的较大值
		depth = int(math.Max(float64(depth), float64(maxDepth(v))))
	}

	return depth + 1
}
