package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var flag bool

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	flag = true
	judge(root, root.Val)
	return flag
}
func judge(root *TreeNode, val int) {
	if root == nil {
		return
	}
	if root.Val != val {
		flag = false
	}
	judge(root.Left, val)
	judge(root.Right, val)
}
