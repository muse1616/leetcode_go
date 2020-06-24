package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var helperFunc func(root *TreeNode)
	helperFunc = func(root *TreeNode) {
		if root==nil{
			return
		}
		//	遍历右子树
		helperFunc(root.Right)
		root.Val = root.Val+ sum
		sum = root.Val
		//	左子树
		helperFunc(root.Left)
	}

	helperFunc(root)
	return root
}

