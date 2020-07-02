package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//先序 [root,(左),(右)]
//中序 [(左),root,(右)]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	//    根节点是先序遍历的第一个节点
	root := &TreeNode{preorder[0], nil, nil}
	//    找到中序中根节点位置 就能确定左右子树的节点数量
	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == preorder[0] {
			break
		}
	}
	//    此时已确定中序中根节点的位置 然后递归确定左子树中根节点位置 右子树中根节点位置 最后接上即可
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}