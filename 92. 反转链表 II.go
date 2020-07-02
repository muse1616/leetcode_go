package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//	头插法
	//	守卫指针放在需要翻转的前一个元素
	//	遍历指针在翻转区间向后移动 每次遍历到一个元素 就删除然后放在守卫指针后面即可
	dummyHead := new(ListNode)
	dummyHead.Next = head
	//守卫在遍历前一个位置
	guardPoint := dummyHead
	travelPoint := dummyHead.Next
	step := 0
	//找到翻转位置
	for step < m-1 {
		guardPoint = guardPoint.Next
		travelPoint = travelPoint.Next
		step++
	}
	//头插 头插次数为n-m
	for i := 0; i < n-m; i++ {
		tmp := travelPoint.Next
		//删除
		travelPoint.Next = travelPoint.Next.Next
		//	加到守护指针前
		// 注意链表不要断开
		tmp.Next = guardPoint.Next
		guardPoint.Next = tmp
	}

	return dummyHead.Next

}
