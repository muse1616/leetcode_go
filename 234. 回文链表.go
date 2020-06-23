package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	for i:=0;i<=len(list)/2;i++{
		if list[i]!= list[len(list)-i]{
			return false
		}
	}
	return true

}
