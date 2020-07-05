package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	//	队列实现层次遍历
	if root == nil {
		return nil
	}
	//层序遍历
	var result [][]int
	//切片模拟队列
	var queue []*TreeNode
	//将根节点压入队列
	queue = append(queue, root)
	var tmpQueue []*TreeNode
	var level []int

	// 依次出列
	for len(queue) >= 0 {
		if len(queue) == 0 {
			//	队列重置 答案重置
			queue = append(queue, tmpQueue...)
			tmpQueue = make([]*TreeNode, 0)
			result = append(result, level)
			level = make([]int, 0)
		}
		//出列元素
		if len(queue) == 0 {
			break
		}
		tmpNode := queue[0]
		queue = queue[1:]
		level = append(level, tmpNode.Val)
		//将左右元素压入临时队列 空的话就不压了
		if tmpNode.Left != nil {
			tmpQueue = append(tmpQueue, tmpNode.Left)
		}
		if tmpNode.Right != nil {
			tmpQueue = append(tmpQueue, tmpNode.Right)
		}
	}
	//答案取层序遍历每个的最后一个即可
	ans := make([]int, 0)
	for i := 0; i < len(result); i++ {
		ans = append(ans, result[i][len(result[i])-1])
	}
	return ans
}
