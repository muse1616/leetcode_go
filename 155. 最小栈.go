package leetcode

type MinStack struct {
	val []int
}

var mS MinStack

/** initialize your data structure here. */
func Constructor() MinStack {
	var val1 []int
	var val2 []int
	c := MinStack{val: val1}
	mS = MinStack{val: val2}
	return c
}

func (this *MinStack) Push(x int) {
	this.val = append(this.val, x)
	if len(mS.val) == 0 || x <= mS.Top() {
		mS.val = append(mS.val, x)
	}

}

func (this *MinStack) Pop() {
	if mS.Top() == this.Top() {
		mS.val = mS.val[:len(mS.val)-1]
	}
	this.val = this.val[:len(this.val)-1]
}

func (this *MinStack) Top() int {
	return this.val[len(this.val)-1]
}

func (this *MinStack) GetMin() int {
	return mS.Top()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
