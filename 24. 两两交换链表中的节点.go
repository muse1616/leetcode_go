package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {


	pre := &ListNode{0, head}
	p := pre

	for pre.Next != nil && pre.Next.Next != nil {
		pre.Next.Next.Next, pre.Next.Next, pre.Next, pre = pre.Next, pre.Next.Next.Next, pre.Next.Next, pre.Next




	}
	return p.Next



}