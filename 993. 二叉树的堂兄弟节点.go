package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	parentHash := make(map[int]int, 10)
	depthHash := make(map[int]int, 10)
	dfs(root, nil, 0, &parentHash, &depthHash)

	return depthHash[x] == depthHash[y] && parentHash[x] != parentHash[y]
}
func dfs(root *TreeNode, parent *TreeNode, depth int, pH *map[int]int, dH *map[int]int) {
	if root == nil {
		return
	}
	if parent != nil {
		(*pH)[root.Val] = parent.Val
	}
	(*dH)[root.Val] = depth
	dfs(root.Left, root, depth+1, pH, dH)
	dfs(root.Right, root, depth+1, pH, dH)
}
