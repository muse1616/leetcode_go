package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return helper(root,root)

}

func helper(rootA,rootB *TreeNode)bool  {
	if rootA==nil && rootB == nil{
		return true
	}
	if rootA == nil || rootB == nil{
		return false
	}
	return rootA.Val == rootB.Val && helper(rootA.Left,rootB.Right) && helper(rootA.Right,rootB.Left)
}
