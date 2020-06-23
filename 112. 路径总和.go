package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, sum int) bool {
	var f bool
	f = false
	if root == nil {
		return false
	}
	helper(root, sum)
	return f
}
func helper(root *TreeNode, sum int) {
	if root.Left != nil {
		helper(root.Left, sum-root.Val)
	}
	if root.Right != nil {
		helper(root.Right, sum-root.Val)
	}
	if root.Left == nil && root.Right == nil {
		fmt.Println(sum, root.Val)
		if sum-root.Val == 0 {
			f = true
		}
	}
	return
}
