package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head.Next
	// 1->2->3->4
	// p1 = 1
	// p2 = 2
	p1 := head
	p2 := tmp
	for p1.Next != nil && p2.Next != nil {
		// 1->3
		p1.Next = p2.Next
		// p1 = 3
		p1 = p1.Next
		// 2->4
		p2.Next = p1.Next
		// p2 = 4
		p2 = p2.Next
		//变为
		// 1->3  2->4
		// p1= 3
		// p2= 4
	}

	p1.Next = tmp
	return head
}
