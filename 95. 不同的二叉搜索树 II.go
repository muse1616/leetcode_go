package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 每一个数字 i 作为根节点 所有的排序就是 1~i-1 作为左子树   i+1~n作为右子树 组合起来就是
// 递归完成
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return helper(1, n)
}
func helper(start int, end int) []*TreeNode {
	ans := make([]*TreeNode, 0)
	// 起始点大于最大数字 空节点
	if start > end {
		ans = append(ans, nil)
		return ans
	}
	// 以start-end中的任意一个节点作为根节点进行缔造
	for i := start; i <= end; i++ {
		//	得到所有左子树
		left := helper(start, i-1)
		//	得到所有右子树
		right := helper(i+1, end)
		//	组合左右
		for _, lTree := range left {
			for _, rTree := range right {
				root := new(TreeNode)
				root.Val = i
				root.Left = lTree
				root.Right = rTree
				ans = append(ans, root)
			}
		}
	}
	return ans
}
