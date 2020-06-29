package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum = 0
	traverse(root, L, R)
	return sum
}
func traverse(root *TreeNode, L int, R int) {
	if root == nil {
		return
	}
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}
	traverse(root.Left, L, R)
	traverse(root.Right, L, R)
}
