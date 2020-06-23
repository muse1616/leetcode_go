package leetcode

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var result []string
	tmp := strconv.Itoa(root.Val)
	helper(root, tmp, &result)
	return result
}

func helper(root *TreeNode, tmp string, result *[]string) {

	if root.Left != nil && root.Right != nil {
		newPath := tmp
		tmp = tmp + "->" + strconv.Itoa(root.Left.Val)
		newPath = newPath + "->" + strconv.Itoa(root.Right.Val)
		helper(root.Right, newPath, result)
		helper(root.Left, tmp, result)
	} else if root.Left != nil && root.Right==nil{
		tmp = tmp + "->" + strconv.Itoa(root.Left.Val)
		helper(root.Left, tmp, result)
	} else if root.Right != nil && root.Left == nil{
		tmp = tmp + "->" + strconv.Itoa(root.Right.Val)
		helper(root.Right, tmp, result)
	} else {
		*result = append(*result,tmp)
	}
}