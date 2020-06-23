package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)

	//根节点入列
	queue = append(queue, root)

	//表示从左边开始 false表示从右边开始
	isLeft := true

	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, l)
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if isLeft {
				tmp[i] = node.Val
			} else {
				tmp[l-i-1] = node.Val
			}
		}
		result = append(result, tmp)
		isLeft = !isLeft
		//清空队列
		queue = queue[l:]
	}
	return result

}
