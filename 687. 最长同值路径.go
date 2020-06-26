package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max = 0
	helper(root, root.Val)
	return max
}
func helper(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, root.Val)
	right := helper(root.Right, root.Val)
	max = int(math.Max(float64(max), float64(left+right)))
	if root.Val == val {
		return int(math.Max(float64(left), float64(right))) + 1
	}
	return 0
}
