package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	length := 1
	tmp := head
	for tmp.Next != nil {
		length++
		tmp = tmp.Next
	}
	length = length/2 + 1
	tmp = head
	for i := 1; i < length; i++ {
		tmp = tmp.Next
	}
	return tmp
}
