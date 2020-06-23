package leetcode

type MyQueue struct {
	val   []int
	front int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	var val []int
	return MyQueue{val: val, front: 0}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.val = append(this.val, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	a := this.val[this.front]
	this.front++
	return a
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.val[this.front]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.val) == this.front {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
