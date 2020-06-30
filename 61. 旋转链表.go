package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	//	双指针
	var pre, last *ListNode
	cur := head
	count := 0
	//链表长度
	for cur != nil {
		count++
		last = cur
		cur = cur.Next
	}
	// last指向最后一个节点

	k = k % count
	//循环到了头结点
	if k == 0 {
		return head
	}
	//	倒数第k个 正数 count -k
	cur = head
	for i := 0; i < count-k; i++ {
		pre = cur
		cur = cur.Next
	}
	// 新的链表的头结点的前一个节点指向空
	pre.Next = nil
	// 尾接头
	last.Next = head
	return cur

}
