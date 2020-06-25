package leetcode

type Node struct {
	Val      int
	Children []*Node
}

//var list []int
//func preorder(root *Node) []int {
//	list = make([]int,0)
//	traverse(root)
//	return list
//}
//func traverse(root *Node){
//	if root!=nil{
//		list = append(list,root.Val)
//		for _,v:=range root.Children{
//			traverse(v)
//		}
//	}
//}
func preorder(root *Node) []int {
	var res []int
	var stack = []*Node{root}
	for 0 < len(stack) {
		for root != nil {
			res = append(res, root.Val) //前序输出
			if 0 == len(root.Children) {
				break
			}
			for i := len(root.Children) - 1; 0 < i; i-- {
				stack = append(stack, root.Children[i]) //入栈
			}
			root = root.Children[0]
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}
