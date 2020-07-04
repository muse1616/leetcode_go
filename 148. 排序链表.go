package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	arr := make([]*ListNode, 0)
	tmp := head
	for tmp != nil {
		arr = append(arr, tmp)
		tmp = tmp.Next
	}

	// 归并排序
	mergeSort(arr, 0, len(arr))

	for i := 0; i < len(arr)-1; i++ {
		arr[i].Next = arr[i+1]
	}
	arr[len(arr)-1].Next = nil
	return arr[0]
}
func mergeSort(arr []*ListNode, begin int, end int) {
	if end-begin > 1 {
		//	一分为二
		middle := begin + (end-begin+1)/2
		//	先排左边
		mergeSort(arr, begin, middle)
		//	再排右边
		mergeSort(arr, middle, end)
		//	合并
		merge(arr, begin, middle, end)
	}

}
func merge(arr []*ListNode, begin int, middle int, end int) {
	//	获取左边数组长度
	leftSize := middle - begin
	//	获取右边数组长度
	rightSize := end - middle
	//	临时数组
	tmpArr := make([]*ListNode, 0, leftSize+rightSize)
	//	合并
	// 左右游标
	i, j := 0, 0
	for i < leftSize && j < rightSize {
		if arr[begin+i].Val < arr[middle+j].Val {
			tmpArr = append(tmpArr, arr[begin+i])
			i++
		} else {
			tmpArr = append(tmpArr, arr[middle+j])
			j++
		}
	}
	//	剩余
	tmpArr = append(tmpArr, arr[begin+i:middle]...)
	tmpArr = append(tmpArr, arr[middle+j:end]...)
	// 赋值给原来数组 释放临时数组空间
	for i := 0; i < leftSize+rightSize; i++ {
		arr[begin+i] = tmpArr[i]
	}

}
