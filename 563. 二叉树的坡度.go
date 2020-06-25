package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
var sum int
func findTilt(root *TreeNode) int {
	sum=0
	traverse(root)
	return sum
}
func traverse(root *TreeNode)int  {
	if root==nil{
		return 0
	}
	left:=traverse(root.Left)
	right:=traverse(root.Right)
	sum+=int(math.Abs(float64(left-right)))
	return root.Val+left+right
}