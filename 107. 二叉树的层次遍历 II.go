package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var queue []*TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	//root进队
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue)
		tmp := make([]int, l)
		for i := 0; i < l; i++ {
			//		依次入列
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
		}
		//	清空队列
		queue = queue[l:]

		//需要在前方插入
		tmpA := make([][]int, 1)
		tmpA = append(tmpA, tmp)
		result = append(tmpA, result...)

	}
	return result
}
