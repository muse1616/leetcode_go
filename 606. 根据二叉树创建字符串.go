package leetcode

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	//左右子树都为空
	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val)
	}
	//左子树为空
	if t.Right == nil {
		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"
	}
	//左右子树都存在
	return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")" + "(" + tree2str(t.Right) + ")"
}
