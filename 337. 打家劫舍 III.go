package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func rob(root *TreeNode)int{
	hash := make(map[*TreeNode]int,10)
	return helper(root,&hash)
}
// 根 儿子 孙子 即 max(根+4*孙子,儿子*2)
func helper(root *TreeNode, hash *map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if _,ok:=(*hash)[root];ok{
		return (*hash)[root]
	}
	sum := root.Val
	if root.Left != nil {
		sum += helper(root.Left.Left,hash) + helper(root.Left.Right,hash)
	}
	if root.Right != nil {
		sum += helper(root.Right.Left,hash) + helper(root.Right.Right,hash)
	}
	result :=max(sum, helper(root.Left,hash)+helper(root.Right,hash))
	(*hash)[root] =result
	return result
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
