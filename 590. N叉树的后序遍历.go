package leetcode

type Node struct {
	Val      int
	Children []*Node
}

var list []int

func postorder(root *Node) []int {
	list = make([]int, 0)
	traverse(root)
	return list
}
func traverse(root *Node) {
	if root == nil {
		return
	}
	for _, v := range root.Children {
		traverse(v)
	}
	list = append(list, root.Val)
}
