package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result int

func diameterOfBinaryTree(root *TreeNode) int {
	result = 0
	depth(root)
	return result
}
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//左子树深度
	left := depth(root.Left)
	//右子树深度
	right := depth(root.Right)
	//已某一点为根节点 左+右的深度就是路径长度-1
	result = int(math.Max(float64(result), float64(left+right)))
	//子树的深度 左子树和右子树深度中较大值
	return int(math.Max(float64(left), float64(right))) + 1
}
