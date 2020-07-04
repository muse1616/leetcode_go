package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//将节点存储到线性切片中 然后前后指针取
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	nodeList := make([]*ListNode, 0)
	tmp := head
	for tmp != nil {
		nodeList = append(nodeList, tmp)
		tmp = tmp.Next
	}
	left, right := 0, len(nodeList)-1
	head = nodeList[0]
	for left < right {
		nodeList[left].Next = nodeList[right]
		left++
		//	偶数节点 左右会提前相遇
		if left == right {
			break
		}
		nodeList[right].Next = nodeList[left]
		right--
	}
	//最后那个节点next指向nil
	nodeList[left].Next = nil
}
