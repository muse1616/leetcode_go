package leetcode

type MyHashSet struct {
	value []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	value := make([]int, 0, 10000)
	return MyHashSet{value}
}

func (this *MyHashSet) Add(key int) {
	for _, v := range this.value {
		if v == key {
			return
		}
	}
	this.value = append(this.value, key)
}

func (this *MyHashSet) Remove(key int) {
	for k, v := range this.value {
		if v == key {
			this.value = append(this.value[:k], this.value[k+1:]...)
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	for _, v := range this.value {
		if v == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
