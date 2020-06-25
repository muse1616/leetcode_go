package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var target int
var result bool
var hash map[int]bool

func findTarget(root *TreeNode, k int) bool {
	result = false
	target = k
	hash = make(map[int]bool)
	helper(root)
	return result
}
func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left)
	if hash[target-root.Val] == true {
		result = true
		return
	}
	hash[root.Val] = true
	helper(root.Right)

}
