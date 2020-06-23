package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tmp := head
	for tmp.Next != nil {
		if tmp.Next.Val == tmp.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}

	}
	return head
}
