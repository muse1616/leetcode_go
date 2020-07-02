package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序 [(左),root,(右)]
// 后序[(左),(右),root]
func buildTree(inorder []int, postorder []int) *TreeNode {
	//递归结束
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	//	根节点是后序遍历的最后一个节点
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	//	找到中序中的root节点 确定左右子树数量
	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == root.Val {
			break
		}
	}
	//	此时index为中序中根的位置
	//	递归连接左右子树即可
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root

}
