package leetcode

type MyHashMap struct {
	key   []int
	value []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	key := make([]int, 0, 10)
	value := make([]int, 0, 10)
	return MyHashMap{key, value}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	for i, v := range this.key {
		if v == key {
			this.value[i] = value
			return
		}
	}
	this.key = append(this.key, key)
	this.value = append(this.value, value)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	for i, v := range this.key {
		if v == key {
			return this.value[i]
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	for i, v := range this.key {
		if v == key {
			this.key = append(this.key[:i],this.key[i+1:]...)
			this.value = append(this.value[:i],this.value[i+1:]...)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
