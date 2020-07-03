package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	//	左右都有 左连右
	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}
	if root.Left != nil && root.Right == nil {
		tmp := root.Next
		for tmp != nil {
			if tmp.Left != nil {
				root.Left.Next = tmp.Left
				break
			} else if tmp.Right != nil {
				root.Left.Next = tmp.Right
				break
			}
			tmp = tmp.Next
		}
	}
	if root.Right != nil {

		tmp := root.Next

		for tmp != nil {
			if tmp.Left != nil {
				root.Right.Next = tmp.Left
				break
			} else if tmp.Right != nil {
				root.Right.Next = tmp.Right
				break
			}
			tmp = tmp.Next
		}
	}
	// 这里一定先连右节点 因为左边取决于右边的next！
	connect(root.Right)
	connect(root.Left)
	return root

}
