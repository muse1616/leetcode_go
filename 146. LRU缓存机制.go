package leetcode

type LRUCache struct {
	capacity int
	hash     map[int]int
	queue    []int
}

// 每次put或get 都将改值置换到队列的最末尾 每次删除队列的第一个元素
func Constructor(capacity int) LRUCache {
	h := make(map[int]int, capacity)
	q := make([]int, 0)
	return LRUCache{capacity: capacity, hash: h, queue: q}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.hash[key]; ok {
		//先删除队列中的该key
		for k, v1 := range this.queue {
			if v1 == key {
				this.queue = append(this.queue[:k], this.queue[k+1:]...)
				break
			}
		}
		//添加到队列末尾
		this.queue = append(this.queue, key)
		//返回value
		return v
	}
	//未找到
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	//	本来就存在 修改即可
	if _, ok := this.hash[key]; ok {
		this.hash[key] = value
		//先删除队列中的该key
		for k, v1 := range this.queue {
			if v1 == key {
				this.queue = append(this.queue[:k], this.queue[k+1:]...)
				break
			}
		}
		//添加到队列末尾
		this.queue = append(this.queue, key)
	} else {
		//	不存在
		//	1.容量未满
		if len(this.hash) < this.capacity {
			//	直接加入
			this.hash[key] = value
			//	加入使用队列
			//添加到队列末尾
			this.queue = append(this.queue, key)
		} else {
			//	2.容量已满 删除队列第一个
			delete(this.hash, this.queue[0])
			this.queue = this.queue[1:]
			this.hash[key] = value
			//添加到队列末尾
			this.queue = append(this.queue, key)
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
