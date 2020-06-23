package leetcode

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	BST := []*TreeNode{}

	for root != nil {
		BST = append(BST, root)
		root = root.Left
	}
	return BSTIterator{
		stack: BST,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	n := len(this.stack)
	cur := this.stack[n-1]
	this.stack = this.stack[:n-1]

	root := cur.Right
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}

	return cur.Val
}


func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
