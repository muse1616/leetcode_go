package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	//	快慢指针
	fast := head
	slow := head
	for {
		if fast.Next == nil {
			return false
		} else {
			if fast.Next != nil && fast.Next.Next != nil {
				fast = fast.Next.Next
				slow = slow.Next
			} else {
				fast = fast.Next
			}
		}
		if fast == slow {
			return true
		}

	}
}
