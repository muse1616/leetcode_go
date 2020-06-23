package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val == sum {
		res++
	}

	res += dfs(root.Left, sum-root.Val)
	res += dfs(root.Right, sum-root.Val)
	return res
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return dfs(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
