package leetcode


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 删除的若是头结点
	// 三指针并行 pre p1=p2+1 p3=p2+n
	// p1和p2间隔n-1 p2到尾时删除p1即可
	pre := head
	p1 := head
	p2 := head
	for i := 0; i < n-1; i++ {
		p2 = p2.Next
	}

	count := 0

	for {

		// 此时删除p1
		if p2.Next == nil {
			// 若p1是头结点
			if p1 == head {
				// 只有一个节点 返回空
				if p1.Next==nil{
					return nil
				}
				// 大于一个节点
				head=p1.Next
				return head
			} else {
				pre.Next = p1.Next
				return head
			}
		}

		if count == 0 {
			p1 = p1.Next
			p2 = p2.Next
			count++
		} else {
			p1 = p1.Next
			p2 = p2.Next
			pre = pre.Next
		}
	}
}
