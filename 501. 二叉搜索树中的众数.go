package leetcode

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
var hash map[int]int
func findMode(root *TreeNode) []int{
	hash = make(map[int]int,10)
	travel(root)
	max:=0
	//找到最大值
	for _,v:=range hash{
		if v>max{
			max = v
		}
	}
	var res []int
	for k,v:=range hash{
		if v == max{
			res = append(res,k)
		}
	}
	return res
}
func travel(root *TreeNode){
	if root==nil{
		return
	}
	if _,ok:=hash[root.Val];ok{
		hash[root.Val]++
	}else {
		hash[root.Val] = 1
	}
	travel(root.Left)
	travel(root.Right)
}
