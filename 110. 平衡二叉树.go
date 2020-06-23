package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return root == nil || (isBalanced(root.Left)&&isBalanced(root.Right)&&math.Abs(heigt(root.Right)-heigt(root.Left))<2)
}
func heigt(root *TreeNode)float64  {
	if root == nil{
		return 0
	}
	return math.Max(heigt(root.Left),heigt(root.Right))+1
}