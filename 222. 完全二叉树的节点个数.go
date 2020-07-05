package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int
func countNodes(root *TreeNode) int {
	ans = 0
	dfs(root)
	return ans
}
func dfs(root *TreeNode){
	if root == nil{
		return
	}
	ans ++
	if root.Left == nil && root.Right == nil{
		return
	}
	if root.Left ==nil || root.Right == nil {
		ans++
		return
	}
	dfs(root.Left)
	dfs(root.Right)

}
