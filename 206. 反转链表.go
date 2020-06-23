package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	var newHead *ListNode

	for head!=nil{
		//旧链表头删
		tmp := head
		head = head.Next
		//新链表头插
		tmp.Next = newHead
		newHead = tmp
	}
	return newHead
}
