package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//每次都挑中间的元素即可
func sortedListToBST(head *ListNode) *TreeNode {
	//	链表换成数组
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	return helper(list)
}
func helper(nums []int) *TreeNode {
	//结束
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{nums[len(nums)/2], helper(nums[:len(nums)/2]), helper(nums[len(nums)/2+1:])}
}
