package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	left, right := &ListNode{}, &ListNode{}
	tmpLeft, tmpRight := left, right
	for head != nil {
		if head.Val < x {
			tmpLeft.Next = head
			tmpLeft = tmpLeft.Next
		} else {
			tmpRight.Next = head
			tmpRight = tmpRight.Next
		}
		head = head.Next
	}
	tmpRight.Next = nil
	// 接上
	tmpLeft.Next = right.Next
	return left.Next

}
