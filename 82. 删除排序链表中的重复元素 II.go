package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newList := new(ListNode)

	result := newList
	//前驱
	var pre *ListNode

	for head != nil {
		//等于前一个节点 或者 等于后一个节点
		if (pre != nil && head.Val == pre.Val) || (head.Next != nil && head.Next.Val == head.Val) {
		} else {
			// 否则加入到新链表中
			node := &ListNode{head.Val, nil}
			newList.Next = node
			newList = newList.Next
		}
		//前驱向后
		pre = head
		head = head.Next
	}

	return result.Next
}
