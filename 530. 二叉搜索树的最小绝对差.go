package leetcode

import "math"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var min int

func getMinimumDifference(root *TreeNode) int {
	//最小值为 当前节点与左节点的最右节点的值 或是 当前节点与右节点的最左节点的值
	if root == nil {
		return 0
	}
	min = math.MaxInt64
	travel(root)
	return min

}
func travel(root *TreeNode) {
	if root == nil {
		return
	}
	var tmpNode *TreeNode
	tmpNode = root
	if tmpNode.Left!=nil {
		tmpNode = tmpNode.Left
		for tmpNode.Right!=nil{
			tmpNode = tmpNode.Right
		}
		if int(math.Abs(float64(root.Val-tmpNode.Val)))<min{
			min = int(math.Abs(float64(root.Val-tmpNode.Val)))
		}
	}
	tmpNode = root
	if tmpNode.Right!=nil {
		tmpNode = tmpNode.Right
		for tmpNode.Left!=nil{
			tmpNode = tmpNode.Left
		}
		if int(math.Abs(float64(root.Val-tmpNode.Val)))<min{
			min = int(math.Abs(float64(root.Val-tmpNode.Val)))
		}
	}
	travel(root.Left)
	travel(root.Right)
}
