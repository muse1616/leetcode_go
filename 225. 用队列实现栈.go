package leetcode

//放过golang吧 无队列栈 切片完成

type MyStack struct {
	val []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	var val []int
	return MyStack{val: val}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.val = append(this.val, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	a := this.val[len(this.val)-1]
	this.val = this.val[:len(this.val)-1]
	return a
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.val[len(this.val)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.val) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
