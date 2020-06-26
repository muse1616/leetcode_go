package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var min1 int
var min2 int

func findSecondMinimumValue(root *TreeNode) int {
	if root.Left == nil || root == nil {
		return -1
	}
	min1 = math.MaxInt64
	min2 = min1
	traverse(root)
	if min2 == math.MaxInt64{
		return -1
	}
	return min2
}
func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	//判断当前值
	if root.Val < min1 {
		min2 = min1
		min1 = root.Val
	}
	if root.Val < min2 && root.Val > min1 {
		min2 = root.Val
	}
	traverse(root.Left)
	traverse(root.Right)
}
