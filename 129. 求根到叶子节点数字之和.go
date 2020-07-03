package leetcode

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var strArr []string

func sumNumbers(root *TreeNode) int {
	strArr = make([]string, 0)
	traverse(root, "")
	result := 0
	for _, v := range strArr {
		n, _ := strconv.Atoi(v)
		result += n
	}
	return result
}

func traverse(root *TreeNode, str string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		strArr = append(strArr, str+strconv.Itoa(root.Val))
		return
	}
	traverse(root.Left, str+strconv.Itoa(root.Val))
	traverse(root.Right, str+strconv.Itoa(root.Val))
}
