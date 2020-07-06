package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result int
var index int
func kthSmallest(root *TreeNode, k int) int {
	result = 0
	index = k
	dfs(root)
	return result
}
func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		index--
		if index == 0 {
			result =  root.Val
			return
		}
		dfs(root.Right)
	}
}
