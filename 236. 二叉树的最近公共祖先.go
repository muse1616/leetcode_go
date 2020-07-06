package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	// 自身是父节点 同时判断在该部分子树中是否包含了p或者q
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// 递归左右子树 此时先到达二叉树底部
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 左子树和右子树同时包含p和q 说明左右子树各一个 这个根节点就是最近祖先
	if left != nil && right != nil {
		return root
	}
	// 左子树不包含pq 说明只用看右子树
	if left == nil {
		return right
	}
	// 因为一定存在 所以右子树不存在时 看左子树
	return left
}
