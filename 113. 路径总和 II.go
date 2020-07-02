package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	result = make([][]int, 0)
	backtrace(root, []int{}, sum)
	return result
}
func backtrace(root *TreeNode, num []int, rest int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if rest == root.Val {
			cpy := make([]int, len(num)+1)
			copy(cpy, append(num,root.Val))
			result = append(result, cpy)
		}
		return
	}
	num = append(num, root.Val)
	backtrace(root.Left, num, rest-root.Val)
	backtrace(root.Right, num, rest-root.Val)
}
