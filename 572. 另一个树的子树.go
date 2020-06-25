package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// t == s / t是s左子树 /t是s右子树
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return isSub(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSub(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if (s == nil && t != nil) || (s != nil && t == nil) || (s.Val != t.Val) {
		return false
	}
	return isSub(s.Left, t.Left) && isSub(s.Right, t.Right)
}


