package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int
func bstToGst(root *TreeNode) *TreeNode {
	sum=0
	helper(root)
	return root
}
func helper(root *TreeNode){
	if root == nil{
		return
	}
//	右子树
	helper(root.Right)
	root.Val = root.Val+sum
	sum = root.Val
//	左子树
	helper(root.Left)

}

