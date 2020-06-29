package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var list1 []int
var list2 []int

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	list1 = make([]int, 10)
	list2 = make([]int, 10)
	traverse(root1, &list1)
	traverse(root2, &list2)
	if len(list1) != len(list2) {
		return false
	}
	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}
func traverse(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*list = append(*list, root.Val)
	}
	traverse(root.Left, list)
	traverse(root.Right, list)
}
