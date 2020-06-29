package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var list []int

func increasingBST(root *TreeNode) *TreeNode {
	list = make([]int, 0)
	traverse(root)
	ans := &TreeNode{
		Val:   list[0],
		Right: nil,
		Left:  nil,
	}
	tmp := ans
	for i := 1; i < len(list); i++ {
		next := &TreeNode{
			Val:   list[i],
			Right: nil,
			Left:  nil,
		}
		tmp.Right = next
		tmp = next

	}
	return ans
}
func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left)
	list = append(list, root.Val)
	traverse(root.Right)
}
