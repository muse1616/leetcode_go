package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func sumOfLeftLeaves(root *TreeNode) int {
	sum = 0
	helper(root,false)
	return sum
}

//f代表标志位 递归遍历时 传入说明其是否是左子树 true代表左子树
func helper(root *TreeNode, f bool) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if f == true {
			sum += root.Val
		}
		return
	}
	helper(root.Left, true)
	helper(root.Right, false)
}
