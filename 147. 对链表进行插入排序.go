package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	// 0 / 1
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	next := head.Next
	for next != nil {
		//已经排序 跳过
		if next.Val >= pre.Val {
			next = next.Next
			pre = pre.Next
			continue
		}
		//此时next.Val < pre.Val 需要寻找插入位置
		nodePre := head
		nodeLast := head.Next
		//此时应该插入到头结点之前的位置
		if next.Val < nodePre.Val {
			pre.Next = next.Next
			next.Next = head
			head = next
			next = pre.Next
			//插入完成
			continue
		}
		//	第二种是插入到nodePre~nodeLast
		for nodePre != pre {
			if nodePre.Val <= next.Val && nodeLast.Val >= next.Val {
				pre.Next = next.Next
				nodePre.Next = next
				next.Next = nodeLast
				next = pre.Next
				break
			}
			nodePre = nodePre.Next
			nodeLast = nodeLast.Next
		}
	}
	return head
}
