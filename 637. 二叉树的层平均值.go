package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//层次遍历
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	var result []float64

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		//	记录当前层数量
		count := len(queue)
		//	总和
		sum := 0

		for i := 0; i < count; i++ {
			sum += queue[i].Val
			//	非空则加入
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		//计算均值加入
		result = append(result, float64(sum)/float64(count))
		//出列
		queue = queue[count:]

	}

	return result
}
