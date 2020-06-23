package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	都在左子树 即已左子树为根节点
//	都在右子树 即已右子树为根节点
//	上述不成立 即找到

	//都在左子树
	if p.Val<root.Val&&q.Val<root.Val{
		return lowestCommonAncestor(root.Left,p,q)
	}
	//都在右子树
	if p.Val>root.Val&&q.Val>root.Val{
		return lowestCommonAncestor(root.Right,p,q)
	}
	return  root
}
