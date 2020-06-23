package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	helper(root, &result)
	return result
}
func helper(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, result)
	*result = append(*result, root.Val)
	helper(root.Right, result)

}
