package leetcode

type NumArray struct {
	array []int
}


func Constructor(nums []int) NumArray {
	return NumArray{nums}
}


func (this *NumArray) SumRange(i int, j int) int {
	result:=0
	for ;i<=j;i++{
		result += this.array[i]
	}
	return result
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */