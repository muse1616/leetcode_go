package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根大于左 小于右
func isValidBST(root *TreeNode) bool {
	return traverse(root,math.MinInt64,math.MaxInt64)
}
func traverse(root *TreeNode,valMin int,valMax int)bool  {
	if root ==nil{
		return true
	}
	// 小于最小值或者大于最大值
	if root.Val<=valMin || root.Val>=valMax{
		return false
	}
	// 左子树 最大值是当前根节点 最小值不变
	// 右子树 最小值是当前根节点 最大值不变
	return traverse(root.Left,valMin,root.Val) && traverse(root.Right,root.Val,valMax)
}
