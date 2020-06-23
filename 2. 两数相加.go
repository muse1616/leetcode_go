package leetcode

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	tmp := result
	//进位
	carry := 0
	for l1 != nil || l2 != nil {
		x,y:=0,0
		// 需要分别判断
		if l1!=nil {
			x=l1.Val
			l1=l1.Next
		}
		if l2!=nil {
			y=l2.Val
			l2=l2.Next
		}
		sum := x + y + carry
		tmp.Next = &ListNode{Val: sum % 10}
		tmp = tmp.Next
		carry = sum / 10
	}
	if carry!=0{
		tmp.Next = &ListNode{Val:carry}
	}
	return result.Next
}
